import * as grpc from '@grpc/grpc-js';
import { ApiQueryServiceMethods } from '../apiquerysvc';
import { apiQueryProto, logger } from '../server';
import { ApiQueryRequest } from '../proto/grpc/subdomain/api/apiquery/v1/ApiQueryRequest';
import { ApiQueryResponse, ApiQueryResponse__Output } from '../proto/grpc/subdomain/api/apiquery/v1/ApiQueryResponse';
import { ApiQueryServiceClient } from '../proto/grpc/subdomain/api/apiquery/v1/ApiQueryService';

const {expect} = require('chai');
import * as mocha from 'mocha';

class MockApiQueryComponent implements ApiQueryServiceMethods {

    async GetSubdomainsByApiQuery(call: grpc.ServerUnaryCall<ApiQueryRequest, ApiQueryResponse>, callback: grpc.sendUnaryData<ApiQueryResponse>) {
        const {target} = call.request;
        if (!target) {
            callback({
                message: "Subdomain not found",
                code: grpc.status.INVALID_ARGUMENT
            }, undefined)
            return
        }

        callback(null, {
            subdomains: ["subdomain-one", "subdomain-two"]
        })
    }
}

function startup(port: string, server: grpc.Server){
    const apiQueryComponent = new MockApiQueryComponent();
    const getSubdomainsByApiQuery = apiQueryComponent.GetSubdomainsByApiQuery
    server.addService(apiQueryProto.ApiQueryService.service, {getSubdomainsByApiQuery});

    server.bindAsync(
        `[::]:${port}`,
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

const port = "50051";

mocha.describe('ApiQueryService Server Test', ()=>{
    let client: ApiQueryServiceClient;
    let server: grpc.Server;

    before(done => {
        server = new grpc.Server();
        logger.info("here")
        startup(port, server);
        client = new apiQueryProto.ApiQueryService(`[::]:${port}`, grpc.credentials.createInsecure());
        done()
    });

    after(done => {
        server.tryShutdown(done);
    });

    it("should return two subdomains as response", done =>{
        const request = {
            target: "vmrw.com"
        };

        client.getSubdomainsByApiQuery(request, (err: grpc.ServiceError | null, resp: ApiQueryResponse__Output | undefined) => {
            expect(resp?.subdomains).to.have.lengthOf(2);
            done()
        });
    })

    it("should return a gRPC status Invalid Argument", done =>{
        const request = {
            target: ""
        };

        client.getSubdomainsByApiQuery(request, (err: grpc.ServiceError | null, resp: ApiQueryResponse__Output | undefined) => {
            const statusR = grpc.status.INVALID_ARGUMENT
            expect(err?.code).to.equal(statusR);
            expect(resp).to.be.undefined;
            done()
        });
    })
})
