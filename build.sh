#!/usr/bin/env bash

rm -rf tsprotos && mkdir tsprotos

protoc -I ./protos/ ./protos/utils.proto --go_out=plugins=grpc:tsprotos