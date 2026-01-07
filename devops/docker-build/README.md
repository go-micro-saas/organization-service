# build

统一端口

* http: 8080
* grpc: 50051

```shell
docker build \
    --build-arg BUILD_FROM_IMAGE=golang:1.24.10 \
    --build-arg RUN_SERVICE_IMAGE=debian:stable-20251117 \
    --build-arg APP_DIR=app \
    --build-arg SERVICE_NAME=org-service \
    --build-arg VERSION=latest \
    -t org-service:latest \
    -f ./devops/docker-build/Dockerfile .
```