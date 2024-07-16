package com.subdomain;
import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.exporter.otlp.trace.OtlpGrpcSpanExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.resources.Resource;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.BatchSpanProcessor;
import io.opentelemetry.semconv.*;

import java.util.concurrent.TimeUnit;

public class Telemetry {

    private static final String SERVICE_NAME = "dnsresolvesvc";
    private static final String SERVICE_VERSION = "0.1.0";

    static OpenTelemetry initOpenTelemetry(String endpoint) {
        OtlpGrpcSpanExporter spanExporter =
            OtlpGrpcSpanExporter.builder()
                .setEndpoint(endpoint)
                .setTimeout(30, TimeUnit.SECONDS)
                .build();

        Resource serviceNameResource =
            Resource.create(Attributes.of(ResourceAttributes.SERVICE_NAME, SERVICE_NAME, ResourceAttributes.SERVICE_VERSION, SERVICE_VERSION));

        // This lets us create a tracer. If absent, Otel will use no-op and fail operation
        SdkTracerProvider tracerProvider =
            SdkTracerProvider.builder()
                .addSpanProcessor(BatchSpanProcessor.builder(spanExporter).build())
                .setResource(Resource.getDefault().merge(serviceNameResource))
                .build();

        OpenTelemetrySdk otel =
            OpenTelemetrySdk.builder()
                .setTracerProvider(tracerProvider)
                .build();

        // Shut SDK down cleanly
        Runtime.getRuntime().addShutdownHook(new Thread(tracerProvider::close));
        return otel;
    }

}
