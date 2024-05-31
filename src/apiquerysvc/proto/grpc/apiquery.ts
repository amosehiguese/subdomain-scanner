import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { ApiQueryServiceClient as _subdomain_api_apiquery_v1_ApiQueryServiceClient, ApiQueryServiceDefinition as _subdomain_api_apiquery_v1_ApiQueryServiceDefinition } from './subdomain/api/apiquery/v1/ApiQueryService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  subdomain: {
    api: {
      apiquery: {
        v1: {
          ApiQueryRequest: MessageTypeDefinition
          ApiQueryResponse: MessageTypeDefinition
          ApiQueryService: SubtypeConstructor<typeof grpc.Client, _subdomain_api_apiquery_v1_ApiQueryServiceClient> & { service: _subdomain_api_apiquery_v1_ApiQueryServiceDefinition }
        }
      }
    }
  }
}

