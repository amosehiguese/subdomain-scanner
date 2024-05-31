const {NodeTracerProvider} = require('@opentelemetry/sdk-trace-node');
const {SimpleSpanProcessor} = require('@opentelemetry/sdk-trace-base');
const {OTLPTraceExporter} = require('@opentelemetry/exporter-otlp-grpc');
const {Resource} = require('@opentelemetry/resources');
const {SEMRESATTRS_SERVICE_NAME} = require('@opentelemetry/semantic-conventions');

export const initTracing = () => {
    const provider = new NodeTracerProvider({
        resource: new Resource({
            [SEMRESATTRS_SERVICE_NAME]: process.env.OTEL_SERVICE_NAME || 'apiquerysvc',
        }),
    });

    const otelEndpoint = process.env.OTEL_ENDPOINT;

    provider.addSpanProcessor(new SimpleSpanProcessor(new OTLPTraceExporter({url: otelEndpoint})));
    provider.reqister();
};