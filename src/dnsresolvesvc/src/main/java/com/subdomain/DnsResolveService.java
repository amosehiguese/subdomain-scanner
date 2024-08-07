package com.subdomain;

import io.grpc.BindableService;
import io.grpc.Server;
import io.grpc.Grpc;
import io.grpc.ServerBuilder;
import io.grpc.ServerInterceptors;
import io.grpc.ServerServiceDefinition;
import io.grpc.StatusRuntimeException;
import io.grpc.InsecureServerCredentials;
import io.grpc.health.v1.HealthCheckResponse.ServingStatus;
import io.grpc.protobuf.services.HealthStatusManager;
import io.grpc.stub.StreamObserver;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.instrumentation.grpc.v1_6.GrpcTelemetry;

import java.util.List;
import java.util.ArrayList;
import java.util.stream.Collectors;
import java.util.concurrent.Executors;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutorService;

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
    private static final int THREAD_POOL_SIZE = 100;

    private Server server;
    private HealthStatusManager healthMgr;
    private static final OpenTelemetry otel = initTracing();

    private static final DnsResolveService service = new DnsResolveService();

    private void start() throws IOException {
        int port = Integer.parseInt(System.getenv().getOrDefault("PORT", "50053"));
        healthMgr = new HealthStatusManager();

        server =
            Grpc.newServerBuilderForPort(port, InsecureServerCredentials.create())
            // .addService(configureServerInterceptor(otel, getResolveServiceImpl()))
            .addService(getResolveServiceImpl())
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
            ExecutorService executor = Executors.newFixedThreadPool(THREAD_POOL_SIZE);
            try {
                List<String> subdomains = request.getHostsList();

                // Fan-out: Distribute tasks across multiple threads
                List<CompletableFuture<String>> futures = subdomains.stream()
                .map(subdomain -> CompletableFuture.supplyAsync(() -> resolves(subdomain) ? subdomain : null, executor))
                .collect(Collectors.toList());

                // Fan-in: Collect results from all threads
                List<String> resolvedSubdomains = futures.stream()
                .map(CompletableFuture::join)  // Wait for each future to complete and get the result
                .filter(subdomain -> subdomain != null) // Filter out unresolved subdomains
                .collect(Collectors.toList());

                ResolveDnsResponse resp = ResolveDnsResponse.newBuilder().addAllSubdomain(resolvedSubdomains).build();
                responseObserver.onNext(resp);
                responseObserver.onCompleted();
            } catch (StatusRuntimeException e) {
                logger.log(Level.WARN, "Failed with status {}", e.getStatus());
                responseObserver.onError(e);
            } finally {
                executor.shutdown();
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

    private static OpenTelemetry initTracing() {
        String tracing_enabled = System.getenv("TRACING_ENABLED");
        if (tracing_enabled != null && tracing_enabled == "1") {
            String endpoint = System.getenv("OTEL_COLLECTOR_ADDR");
            if (endpoint == null) {
                endpoint = "http://jaeger-otel.jaeger.svc.cluster.local:14278/api/traces";
            }
            OpenTelemetry otele = Telemetry.initOpenTelemetry(endpoint);
            logger.info("Tracing enabled");
            return otele;
        }
        logger.info("Tracing disabled");
        return null;
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        // Start the RPC Server
        logger.info("DnsResolve Service starting...");
        final DnsResolveService service = DnsResolveService.getInstance();

        service.start();
        service.blockUntilShutdown();
    }
}
