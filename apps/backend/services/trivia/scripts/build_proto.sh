SCRIPT_DIR=$(dirname "$0")

PROTO_DIR="$SCRIPT_DIR/../../../protobufs"
OUTPUT_DIR="$SCRIPT_DIR/../api/grpc/protobufs"

protoc --proto_path="$PROTO_DIR" \
    --go_out="$OUTPUT_DIR" --go_opt=paths=source_relative \
    --go-grpc_out="$OUTPUT_DIR" --go-grpc_opt=paths=source_relative \
    "$PROTO_DIR"/trivia.proto