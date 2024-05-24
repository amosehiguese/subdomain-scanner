const pino = require('pino');

export const logger = pino({
    name: 'apiquerysvc-server',
    messageKey: 'message',
    formatters: {
        level: (logLevelString, logLevelNum) => {
            return {severity: logLevelString}
        }
    }
});

// Register GRPC OTEL Instrumentation for trace propagation
const {GrpcInstrumentation} = require('@opentelemetry/instrumentation-grpc');
const {registerInstrumentations} = require('@opentelemetry/instrumentation');

registerInstrumentations({
    instrumentations: [new GrpcInstrumentation()]
});

const {initTracing} = require('./telemetry');

if (process.env.ENABLE_TRACING == "1") {
    logger.info("Tracing enabled.")
    initTracing();
} else {
    logger.info("Tracing disabled.")
}

const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

// Helper function that loads a protobuf file. 
const _loadProto = (path) => {
    const packageDefinition = protoLoader.loadSync(
        path,
        {
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true
        }
    );

    return grpc.loadPackageDefinition(packageDefinition);
}

const MAIN_PROTO_PATH = path.join(__dirname, './proto/api/v1/apiquery.proto');
const HEALTH_PROTO_PATH = path.join(__dirname, './proto/grpc/health/v1/health.proto');

const PORT = process.env.PORT;

const apiQueryProto = _loadProto(MAIN_PROTO_PATH).subdomain.api.apiquery.v1;
const health = _loadProto(HEALTH_PROTO_PATH).grpc.health.v1;

const check = (call, callback) => {
    callback(null, { status: 'SERVING'});
}

// Starts a RPC server that receives requests
const main = () => {
    logger.info(`Starting  gRPC server on port ${PORT}...`);
}


