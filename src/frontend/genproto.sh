#!/bin/bash -eu

protodir=./genproto

protoc --go_out=./genproto --go-grpc_out=./genproto --proto_path=$protodir $protodir/subdomain.proto