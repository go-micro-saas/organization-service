# deploy

实际部署：docker-compose.yaml不映射端口，统一使用`Dockerfile`.`EXPOSE`的端口

> 构建文件： [Dockerfile](../docker-build/Dockerfile)

* http: 8080
* grpc: 50051

```shell
# general config
go run ./app/org-service/cmd/store-configuration/... \
  -conf=./app/org-service/configs \
  -source_dir ./devops/docker-deploy/general-configs \
  -store_dir go-micro-saas/general-configs/testing

# service config
go run ./app/org-service/cmd/store-configuration/... \
  -conf=./app/org-service/configs \
  -source_dir ./devops/docker-deploy/service-configs \
  -store_dir go-micro-saas/org-service/testing/latest
  
# database migration
make run-database-migration
  
# run
docker run -it --rm \
  --entrypoint /bin/bash \
  -v $PWD/devops/docker-deploy/configs:/myworkspace/golang/src/workspace/configs \
  org-service:latest
```