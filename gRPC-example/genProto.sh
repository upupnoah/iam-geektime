#! /bin/zsh

# shellcheck disable=SC2120
function genProto {
  DOMAIN=$1
  PROTO_PATH=./${DOMAIN}
  GO_OUT_PATH=./${DOMAIN}
  mkdir -p "${GO_OUT_PATH}"
  protoc -I="${PROTO_PATH}" --go_out=paths=source_relative:"${GO_OUT_PATH}" --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:"${GO_OUT_PATH}" "${DOMAIN}".proto
}

genProto "$1"