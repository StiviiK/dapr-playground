#!/bin/bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
docker build -t dapr0acr0prexqb.azurecr.io/actor-go-client:$1 -f $DIR/client/Dockerfile $DIR
docker build -t dapr0acr0prexqb.azurecr.io/actor-go-caller:$1 -f $DIR/caller/Dockerfile $DIR