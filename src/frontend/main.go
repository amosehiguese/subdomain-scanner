package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/joho/godotenv/autoload"
)

const (
	port = "8080"
)

var encoderCfg = zapcore.EncoderConfig{
	MessageKey: "msg",
	NameKey:    "name",

	LevelKey:    "level",
	EncodeLevel: zapcore.LowercaseLevelEncoder,

	TimeKey:    "time",
	EncodeTime: zapcore.ISO8601TimeEncoder,
}

type frontendServer struct {
	apiQuerySvcAddr string
	apiQuerySvcConn *grpc.ClientConn

	bruteForceSvcAddr string
	bruteForceSvcConn *grpc.ClientConn

	dnsResolveSvcAddr string
	dnsResolveSvcConn *grpc.ClientConn

	portScanSvcAddr string
	portScanSvcConn *grpc.ClientConn

	otelCollectorAddr string
	otelCollectorConn *grpc.ClientConn

	// aiAssistantSvcAddr string
}

func main() {
	ctx := context.Background()
	zapLog := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		),
		zap.Fields(
			zap.String("version", runtime.Version()),
		),
	)
	defer func() { _ = zapLog.Sync() }()

	svc := new(frontendServer)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{}, propagation.Baggage{},
		),
	)

	if os.Getenv("TRACING_ENABLED") == "1" {
		zapLog.Info("Tracing enabled.")
		initTracing(zapLog, ctx, svc)
	} else {
		zapLog.Info("Tracing disabled.")
	}

	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}

	addr := os.Getenv("LISTEN_ADDR")
	mustMapEnv(&svc.apiQuerySvcAddr, "APIQUERY_SERVICE_ADDR")
	mustMapEnv(&svc.bruteForceSvcAddr, "BRUTE_FORCE_SERVICE_ADDR")
	mustMapEnv(&svc.dnsResolveSvcAddr, "DNS_RESOLVE_SERVER_ADDR")
	mustMapEnv(&svc.portScanSvcAddr, "PORT_SCAN_SERVICE_ADDR")
	// mustMapEnv(&svc.cyberAssistantSvcAddr, "CYBER_ASSISTANT_SERVICE_ADDR")

	mustConnGRPC(ctx, &svc.apiQuerySvcConn, svc.apiQuerySvcAddr)
	mustConnGRPC(ctx, &svc.bruteForceSvcConn, svc.bruteForceSvcAddr)
	mustConnGRPC(ctx, &svc.dnsResolveSvcConn, svc.dnsResolveSvcAddr)
	mustConnGRPC(ctx, &svc.portScanSvcConn, svc.portScanSvcAddr)
	// mustConnGRPC(ctx, &svc.cyberAssistantSvcAddr, svc.cyberAssistantSvcAddr)

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/api/v1/scan/{domain}", svc.scanHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })

	var handler http.Handler = r
	handler = &logHandler{log: zapLog, next: handler}
	handler = otelhttp.NewHandler(handler, "frontend")

	zapLog.Sugar().Infof("starting server on " + addr + ":" + srvPort)
	zapLog.Sugar().Fatal(http.ListenAndServe(addr+":"+srvPort, handler))
}

func initTracing(zapLog *zap.Logger, ctx context.Context, svc *frontendServer) (*sdktrace.TracerProvider, error) {
	mustMapEnv(&svc.otelCollectorAddr, "OTEL_COLLECTOR_ADDR")
	mustConnGRPC(ctx, &svc.otelCollectorConn, svc.otelCollectorAddr)
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithGRPCConn(svc.otelCollectorConn),
	)
	if err != nil {
		zapLog.Sugar().Warnf("warn: Failed to create trace exporter: %v", err)
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(traceProvider)

	return traceProvider, err
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}

	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	_, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithUnaryInterceptor(otelgrpc.NewClientHandler),
		// grpc.WithStreamInterceptor(otelgrpc.NewClientHandler),
	)
	*conn, err = grpc.NewClient(addr, opts...)
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
