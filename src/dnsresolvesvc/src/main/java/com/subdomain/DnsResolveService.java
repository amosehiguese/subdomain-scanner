package com.subdomain;

import io.grpc.BindableService;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.ServerInterceptors;
import io.grpc.ServerServiceDefinition;
import io.grpc.StatusRuntimeException;
import io.grpc.health.v1.HealthCheckResponse.ServingStatus;
import io.grpc.protobuf.services.HealthStatusManager;
import io.grpc.stub.StreamObserver;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.instrumentation.grpc.v1_6.GrpcTelemetry;

import java.util.List;
import java.util.ArrayList;
import java.io.IOException;
import java.net.InetAddress;
import java.net.UnknownHostException;

import com.subdomain.Resolvedns.ResolveDnsRequest;
import com.subdomain.Resolvedns.ResolveDnsResponse;
import com.subdomain.Resolvedns.ResolveDnsResponseOrBuilder;
import com.subdomain.ResolveDnsServiceGrpc.ResolveDnsServiceImplBase;

import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;


public class DnsResolveService {
    private static final Logger logger = LogManager.getLogger(DnsResolveService.class);

    private Server server;
    private HealthStatusManager healthMgr;

    private static OpenTelemetry otel;

    private static final DnsResolveService service = new DnsResolveService();

    private void start() throws IOException {
        int port = Integer.parseInt(System.getenv().getOrDefault("PORT", "50053"));
        healthMgr = new HealthStatusManager();

        server =
            ServerBuilder.forPort(port)
            // .addService(configureServerInterceptor(otel, getResolveServiceImpl()))
            .addService(new DnsResolveServiceImpl())
            .addService(healthMgr.getHealthService())
            .build()
            .start();

        logger.info("DnsResolve Service started, listening on " + port);
        Runtime.getRuntime()
            .addShutdownHook(
                new Thread(
                    () -> {
                        System.err.println(
                            "shutting down gRPC DnsResolve server"
                            );
                        DnsResolveService.this.stop();
                        System.err.println("server shut down");
                    }
                )
            );
        healthMgr.setStatus("", ServingStatus.SERVING);
    }

    // Ensures that all gRPC server requests are automatically traced
    ServerServiceDefinition configureServerInterceptor(
        OpenTelemetry otel,
        BindableService bindableService
    ) {
        GrpcTelemetry grpcTelemetry = GrpcTelemetry.create(otel);
        return ServerInterceptors.intercept(bindableService, grpcTelemetry.newServerInterceptor());
    }

    private void stop() {
        if (server != null){
            healthMgr.clearStatus("");
            server.shutdown();
        }
    }

    public static DnsResolveServiceImpl getResolveServiceImpl(){
        return new DnsResolveServiceImpl();
    }

    private static class DnsResolveServiceImpl extends ResolveDnsServiceImplBase {
        @Override
        public void resolveDns(ResolveDnsRequest request, StreamObserver<ResolveDnsResponse> responseObserver) {
            logger.info("received " + request.getHostsCount() + " new dns to resolve");
            try {
                List<String> result = new ArrayList<>();
                for (String host: request.getHostsList()){
                    if (resolves(host)) {
                        result.add(host);
                    }
                }
                ResolveDnsResponse resp = ResolveDnsResponse.newBuilder().addAllSubdomain(result).build();
                responseObserver.onNext(resp);
                responseObserver.onCompleted();
            } catch (StatusRuntimeException e) {
                logger.log(Level.WARN, "Failed with status {}", e.getStatus());
                responseObserver.onError(e);
            }
        }

        public static boolean resolves(String host) {
            try {
                InetAddress address = InetAddress.getByName(host);
                logger.info(host + " resolved successfully with address " + address);
                return true;

            } catch (UnknownHostException e) {
                logger.info("unable to resolve " + host);
                return false;
            }
        }
    }

    private static DnsResolveService getInstance() {
        return service;
    }

    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    private static void initTracing() {
        String tracing_enabled = System.getenv("tracing_enabled");
        if (tracing_enabled != null && tracing_enabled == "true") {
            String endpoint = System.getenv("OTEL_ENDPOINT");
            otel = Telemetry.initOpenTelemetry(endpoint);
            logger.info("Tracing enabled");
            return;
        }
        logger.info("Tracing disabled");
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        new Thread(
            () -> {
                initTracing();
            }
        )
        .start();

        // Start the RPC Server
        logger.info("DnsResolve Service starting...");
        final DnsResolveService service = DnsResolveService.getInstance();
        service.start();
        service.blockUntilShutdown();
    }
}
