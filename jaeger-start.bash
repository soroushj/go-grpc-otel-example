#!/bin/bash

set -e

docker run -d --rm --name jaeger \
  -p 14268:14268 \
  -p 16686:16686 \
  jaegertracing/all-in-one:1

UI_URL='http://localhost:16686/'
echo "Jaeger UI available at $UI_URL"

# open jaeger ui in the default browser
if command -v open &> /dev/null
then
  open "$UI_URL" &> /dev/null
elif command -v xdg-open &> /dev/null
then
  xdg-open "$UI_URL" &> /dev/null
fi
