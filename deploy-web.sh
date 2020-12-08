#!/bin/bash

export SHA=$(git rev-parse HEAD)

docker build -t happilymarrieddadudemy/node-web-app:latest \
    -f ./simple-resources/frontend/Dockerfile \
    ./simple-resources/frontend

docker push happilymarrieddadudemy/node-web-app:latest

# sleep 10

# kubectl set image deployments/frontend-deployment frontend=happilymarrieddadudemy/node-web-app:latest
