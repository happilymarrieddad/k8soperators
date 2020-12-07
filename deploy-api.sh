#!/bin/bash

export SHA=$(git rev-parse HEAD)

docker build -t happilymarrieddadudemy/node-web-server:${SHA} \
    -f ./simple-resources/api/Dockerfile \
    ./simple-resources/api

docker push happilymarrieddadudemy/node-web-server:${SHA}

# sleep 10

# kubectl set image deployments/api-deployment api=happilymarrieddadudemy/node-web-server:${SHA}
