#!/bin/bash

export JAEGER_URL=http://localhost:14268/api/traces

go run ./client/cmd/client localhost:9090
