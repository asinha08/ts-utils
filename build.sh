#!/usr/bin/env bash

rm -rf tsprotos && mkdir tsprotos

#protoc -I=protos --go_out=tsprotos protos/utils.proto
protoc --proto_path=protos/ --go_out=tsprotos --go_opt=paths=source_relative protos/utils.proto