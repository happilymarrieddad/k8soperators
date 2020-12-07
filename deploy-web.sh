#!/bin/bash

export SHA=$(git rev-parse HEAD)

docker build -t happilymarrieddadudemy/node-web-app:${SHA} \
    -f ./simple-resources/frontend/Dockerfile \
    ./simple-resources/frontend

docker push happilymarrieddadudemy/node-web-app:${SHA}

# sleep 10

# kubectl set image deployments/frontend-deployment frontend=happilymarrieddadudemy/node-web-app:${SHA}
