import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { ApiQueryServiceClient as _subdomain_ApiQueryServiceClient, ApiQueryServiceDefinition as _subdomain_ApiQueryServiceDefinition } from './subdomain/ApiQueryService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  subdomain: {
    ApiQueryRequest: MessageTypeDefinition
    ApiQueryResponse: MessageTypeDefinition
    ApiQueryService: SubtypeConstructor<typeof grpc.Client, _subdomain_ApiQueryServiceClient> & { service: _subdomain_ApiQueryServiceDefinition }
  }
}

