#!/bin/bash

export JAEGER_URL=http://localhost:14268/api/traces

go run ./server/cmd/server :9090
