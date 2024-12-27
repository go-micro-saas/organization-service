.PHONY: store-configuration
# run :-->: run database-migration
store-configuration:
	go run ./app/org-service/cmd/store-configuration/... -conf=./app/org-service/configs

.PHONY: run-database-migration
# run :-->: run database-migration
run-database-migration:
	go run ./app/org-service/cmd/database-migration/... -conf=./app/org-service/configs

.PHONY: run-org-service
# run service :-->: run org-service
run-org-service:
	go run ./app/org-service/cmd/org-service/... -conf=./app/org-service/configs

.PHONY: org-org-service
# org service :-->: org org-service
org-org-service:
	curl http://127.0.0.1:9991/api/v1/testdata/get?message=hello

.PHONY: run-service
# run service :-->: run org-service
run-service:
	#@$(MAKE) run-org-service
	go run ./app/org-service/cmd/org-service/... -conf=./app/org-service/configs

.PHONY: org-service
# org service :-->: org org-service
org-service:
	curl http://127.0.0.1:9991/api/v1/testdata/get?message=hello

