#!/bin/bash
set -e

echo "Generating Go…"
protoc --go_out=. --go-grpc_out=. \
  messages.proto permission.proto brokers.proto

echo "Generating Python…"
python3 -m grpc_tools.protoc \
  -I . \
  --python_out=xpb \
  --grpc_python_out=xpb \
  messages.proto permission.proto

echo "Done!"
