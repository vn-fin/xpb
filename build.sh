#!/bin/bash
set -e

echo "Generating Go…"
protoc --go_out=. --go-grpc_out=. \
  messages.proto permission.proto

echo "Generating Python…"
python3 -m grpc_tools.protoc \
  -I . \
  --python_out=py_xpb \
  --grpc_python_out=py_xpb \
  messages.proto permission.proto

echo "Done!"
