#!/bin/bash
#set -x

cd "$(dirname "$0")"
WASM_FILE="wasm_red.wasm"
BINDGEN_DIR="bindout"

init-rust() {
    rustup target add wasm32-unknown-unknown --toolchain nightly
    rustup update
    rustup default nightly
    cargo +nightly install wasm-bindgen-cli
}

init-protoc() {
    # windows setup
    # https://github.com/protocolbuffers/protobuf/issues/2698
    echo "install protoc first"
}

proto-doc() {
    # go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
    # windows platform
    GOOGLE_INCUDE_PATH=C:/go/src
    PROTOS=../../../proto/stu3/*.proto
    protoc --doc_out=./doc --doc_opt=html,proto-fhir.html --proto_path=../../../ --proto_path=${GOOGLE_INCUDE_PATH} ${PROTOS}
    protoc --doc_out=./doc --doc_opt=markdown,proto-fhir.md --proto_path=../../../ --proto_path=${GOOGLE_INCUDE_PATH} ${PROTOS}
}

server() {
    # wasm mimetype is not supported by the lite-server command
    npm run serve
}

build(){
    mkdir -p ${BINDGEN_DIR}
    cargo  +nightly build --target wasm32-unknown-unknown
    wasm-bindgen target/wasm32-unknown-unknown/debug/${WASM_FILE} --out-dir ./${BINDGEN_DIR}
}

case "$1" in 
    init)   init-rust ;;
    build)   build ;;
    doc)   proto-doc ;;
    *) echo "usage: $0 start|stop|restart" >&2
       exit 1
       ;;
esac
