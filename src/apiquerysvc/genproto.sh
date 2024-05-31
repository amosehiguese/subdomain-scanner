#!bin/bash

# Generate the typescript files from the proto files
yarn proto-loader-gen-types --grpcLib=@grpc/grpc-js --outDir=proto/grpc proto/api/v1/*.proto proto/grpc/health/v1/*.proto

