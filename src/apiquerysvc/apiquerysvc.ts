import * as grpc from '@grpc/grpc-js';
import { ApiQueryRequest } from './proto/grpc/subdomain/api/apiquery/v1/ApiQueryRequest';
import { ApiQueryResponse } from './proto/grpc/subdomain/api/apiquery/v1/ApiQueryResponse';

const {logger} = require('./server');

type Response = {
    subdomains: string[],
}

type ApiResponse = {
  version: string;
  app: string;
  host: string;
  response_time_sec: string;
  status: string;
  result: {
      host: string;
      resolved_ip: string;
      issued_to: string;
      issued_o: string | null;
      issuer_c: string;
      issuer_o: string;
      issuer_ou: string | null;
      issuer_cn: string;
      cert_sn: string;
      cert_sha1: string;
      cert_alg: string;
      cert_ver: number;
      cert_sans: string;
      cert_exp: boolean;
      cert_valid: boolean;
      valid_from: string;
      valid_till: string;
      validity_days: number;
      days_left: number;
      valid_days_to_expire: number;
      hsts_header_enabled: boolean;
  };
};

export interface ApiQueryServiceMethods {
  GetSubdomainsByApiQuery(call: grpc.ServerUnaryCall<ApiQueryRequest, ApiQueryResponse>, callback: grpc.sendUnaryData<ApiQueryResponse>): void;
}

export class ApiQueryComponent implements ApiQueryServiceMethods {
    public async GetSubdomainsByApiQuery(call: grpc.ServerUnaryCall<ApiQueryRequest, ApiQueryResponse>, callback: grpc.sendUnaryData<ApiQueryResponse>) {
      const {target} = call.request;
      try {
        if (!target) {
            callback({
                message: `Subdomains not found for ${target}`,
                code: grpc.status.INVALID_ARGUMENT
            }, null);
        }

        const url = `https://ssl-checker.io/api/v1/check/${target}`;

        const data  = await fetchData<ApiResponse>(url);
        if (data === null) {
          logger.error(`Failed to get subdomains for ${target}`)
          callback({
            message: `Failed to get subdomains for ${target}`,
            code: grpc.status.INVALID_ARGUMENT
          })
          return
        }
        const subdomains = new Set<string>(extractSubdomains(data as ApiResponse))
        logger.info(`Done extracting subdomains for ${target}`)

        const response: Response = {
            subdomains: [...subdomains]
        }
        logger.info("Sending response...")
        callback(null, response);
      } catch (error) {
        logger.error(`Failed to get subdomains for ${target}`)
        callback({
          message: `Failed to get subdomains for ${target}`,
          code: grpc.status.INTERNAL
        })
      }

    }
}


const extractSubdomains = (data: ApiResponse): string[] => {
    if (!data) {
      logger.info("No data received")
      return []
    }
    try {
      logger.info("Extracting subdomains from data")

      const sans: string = data.result.cert_sans;
      const dns = sans.split(";").map(entry => entry.trim());

      const subdomains = dns
        .filter(entry => entry.startsWith("DNS:"))
        .map(entry => entry.replace("DNS:", ""))

      return subdomains
    } catch (error) {
      logger.errorr("An errror occurred trying to extract subdomain", error)
      return []
    }

}

const fetchData = async <T>(url: string): Promise<T|null> => {
    try {
        const resp = await fetch(url);
        if (!resp.ok) {
            return null
        }
        const response: T = await resp.json();
        return response
    } catch (error) {
        logger.error(error, "Failed to fetch data")
        return null
    }
}
