#!/bin/bash
#set -x

cd "$(dirname "$0")"

init() {
    go get -u github.com/golang/protobuf/protoc-gen-go
}

gen() {
    PROTO_DIR=../../../proto/stu3/
    PATH_INCLUDE="-I../../../ -I../../protos"
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/datatypes.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/codes.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/annotations.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/resources.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/extensions.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/google_extensions.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/metadatatypes.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/profile_config.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/profiles.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/uscore_codes.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/uscore.proto
    protoc ${PATH_INCLUDE} --go_out=./out ${PROTO_DIR}/version_config.proto
}

case "$1" in 
    init)   init ;;
    gen)   gen ;;
    *) echo "usage: $0 init|gen" >&2
       exit 1
       ;;
esac
