import * as grpc from '@grpc/grpc-js';
import { apiQueryProto } from './server';
import { ApiQueryResponse__Output } from './proto/grpc/subdomain/api/apiquery/v1/ApiQueryResponse';

const PORT = process.env.PORT || 50051;

const client = new apiQueryProto.ApiQueryService(`[::]:${PORT}`, grpc.credentials.createInsecure());

const getSubdomains = () => {
    const request = {
        target: "lentarex.com"
    };

    client.GetSubdomainsByApiQuery(request, (err: grpc.ServiceError | null, resp: ApiQueryResponse__Output | undefined) => {
        console.log(resp?.subdomains)
    });
}

getSubdomains();
