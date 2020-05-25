#!/bin/bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $DIR

docker build -t dapr0acr0prexqb.azurecr.io/actor-go-client:$1 -f client/Dockerfile .
docker build -t dapr0acr0prexqb.azurecr.io/actor-go-caller:$1 -f caller/Dockerfile .