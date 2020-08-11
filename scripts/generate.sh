#!/bin/bash -e


SOURCE_DIR=/app/service/v1

protoc -I=${SOURCE_DIR} \
    --go_out=${SOURCE_DIR} \
    ${SOURCE_DIR}/sum.proto

echo "Done"
