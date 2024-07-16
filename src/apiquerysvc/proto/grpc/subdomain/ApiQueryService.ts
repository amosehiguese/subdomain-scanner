// Original file: proto/api/v1/apiquery.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { ApiQueryRequest as _subdomain_ApiQueryRequest, ApiQueryRequest__Output as _subdomain_ApiQueryRequest__Output } from '../subdomain/ApiQueryRequest';
import type { ApiQueryResponse as _subdomain_ApiQueryResponse, ApiQueryResponse__Output as _subdomain_ApiQueryResponse__Output } from '../subdomain/ApiQueryResponse';

export interface ApiQueryServiceClient extends grpc.Client {
  GetSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  GetSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  GetSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  GetSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  getSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  getSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  getSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  getSubdomainsByApiQuery(argument: _subdomain_ApiQueryRequest, callback: grpc.requestCallback<_subdomain_ApiQueryResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface ApiQueryServiceHandlers extends grpc.UntypedServiceImplementation {
  GetSubdomainsByApiQuery: grpc.handleUnaryCall<_subdomain_ApiQueryRequest__Output, _subdomain_ApiQueryResponse>;
  
}

export interface ApiQueryServiceDefinition extends grpc.ServiceDefinition {
  GetSubdomainsByApiQuery: MethodDefinition<_subdomain_ApiQueryRequest, _subdomain_ApiQueryResponse, _subdomain_ApiQueryRequest__Output, _subdomain_ApiQueryResponse__Output>
}
