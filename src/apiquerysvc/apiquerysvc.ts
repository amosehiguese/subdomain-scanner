import * as grpc from '@grpc/grpc-js';
import { ApiQueryRequest } from './proto/grpc/subdomain/api/apiquery/v1/ApiQueryRequest';
import { ApiQueryResponse } from './proto/grpc/subdomain/api/apiquery/v1/ApiQueryResponse';
import { ApiQueryServiceDefinition, ApiQueryServiceHandlers } from './proto/grpc/subdomain/api/apiquery/v1/ApiQueryService';

const {logger} = require('./server');

type Response = {
    subdomains: string[],
}

type crtshResponse = {
    issuer_ca_id: number,
    issuer_name: string,
    common_name: string,
    name_value: string,
    id: number,
    entry_timestamp: string,
    not_before: string
    not_after: string,
    serial_number: string,
    result_count: number
}

export interface ApiQueryServiceMethods {
    getSubdomainsByApiQuery(call: grpc.ServerUnaryCall<ApiQueryRequest, ApiQueryResponse>, callback: grpc.sendUnaryData<ApiQueryResponse>): void;
}

export class ApiQueryComponent implements ApiQueryServiceMethods {
    public async getSubdomainsByApiQuery(call: grpc.ServerUnaryCall<ApiQueryRequest, ApiQueryResponse>, callback: grpc.sendUnaryData<ApiQueryResponse>) {
        const {target} = call.request;
        if (!target) {
            callback({
                message: "Subdomain not found",
                code: grpc.status.INVALID_ARGUMENT
            }, null);
            return
        }

        const url = `https://crt.sh/?q=%25.${target}&output=json`;
        
        const data  = await fetchData<crtshResponse>(url);
        const subdomains = new Set<string>(data.map((d: crtshResponse) => d.common_name))
        
        const response: Response = {
            subdomains: [...subdomains]
        }
        callback(null, response);
    }
}


const fetchData = async <T>(url: string): Promise<T[]> => {
    try {
        const resp = await fetch(url);
        if (!resp.ok) {
            return []
        }
        const response: T[] = await resp.json();
        return response
    } catch (error) {
        logger.error(error, "Failed to fetch data")
        return []
    }
}