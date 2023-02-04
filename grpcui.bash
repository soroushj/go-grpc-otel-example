#!/bin/bash

set -e

# install grpcui if not already installed
if ! command -v grpcui &> /dev/null
then
  go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
fi

grpcui -plaintext -port 43941 -proto notes/notes.proto localhost:9090
