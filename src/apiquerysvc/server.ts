import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
const pino = require('pino');

export const logger = pino({
    name: 'apiquerysvc-server',
    messageKey: 'message',
    formatters: {
        level: (logLevelString: string, logLevelNum: string) => {
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

import { ApiQueryComponent } from './apiquerysvc';
import {ProtoGrpcType} from './proto/grpc/apiquery';
import {ProtoGrpcType as HealthProtoGrpcType} from './proto/grpc/health';
import { HealthCheckRequest } from './proto/grpc/grpc/health/v1/HealthCheckRequest';
import { HealthCheckResponse} from './proto/grpc/grpc/health/v1/HealthCheckResponse';

type GrpcType = ProtoGrpcType | HealthProtoGrpcType

// Helper function that loads a protobuf file. 
const _loadProto = (path: string): GrpcType => {
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

    return (grpc.loadPackageDefinition(packageDefinition) as unknown) as GrpcType;
}

const MAIN_PROTO_PATH = path.join(__dirname, './proto/api/v1/apiquery.proto');
const HEALTH_PROTO_PATH = path.join(__dirname, './proto/grpc/health/v1/health.proto');

const PORT = process.env.PORT || 7000;

export const apiQueryProto = (_loadProto(MAIN_PROTO_PATH) as ProtoGrpcType).subdomain.api.apiquery.v1;
const health = (_loadProto(HEALTH_PROTO_PATH) as HealthProtoGrpcType).grpc.health.v1;

const check = (call:grpc.ServerUnaryCall<HealthCheckRequest, HealthCheckResponse>, callback:grpc.sendUnaryData<HealthCheckResponse>): void => {
    callback(null, { status: 'SERVING'});
}

// Starts a RPC server that receives requests
const main = () => {
    logger.info(`Starting gRPC server on port ${PORT}...`);
    const server = _getServer();

    server.bindAsync(
        `[::]:${PORT}`,
        grpc.ServerCredentials.createInsecure(),
        (err, port) => {
            if (err) {
                logger.error(`Got an error starting server -> ${err}`);
                return 
            }
            logger.info(`ApiQuerySvc gRPC server started on port ${port}`);
        }
    )
}

const _getServer = () => {
    const server = new grpc.Server();
    const apiQueryComponent = new ApiQueryComponent();
    const getSubdomainsByApiQuery = apiQueryComponent.getSubdomainsByApiQuery
    server.addService(apiQueryProto.ApiQueryService.service, {getSubdomainsByApiQuery});
    server.addService(health.Health.service, {check});

    return server;
}

main();


