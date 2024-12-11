#!/bin/bash

# Get the directory of the current script
SCRIPT_DIR=$(dirname "$0")

PROTO_PATH="$SCRIPT_DIR/../../protobuf-schemas/*.proto"

yarn proto-loader-gen-types \
  --grpcLib=@grpc/grpc-js \
  --outDir=src/grpc/protobufs \
  $PROTO_PATH
