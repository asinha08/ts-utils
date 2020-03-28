#!/usr/bin/env bash

protoc -I ./protos --js_out=import_style=commonjs,binary:../ts-utils-messages/src/generatedCode utils.proto