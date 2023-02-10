SHELL:=/bin/bash

db_init:
	docker compose -f docker-compose.db.yaml up -d

db_migrate:
	docker compose -f docker-compose.migration.yaml build
	docker run --rm naufalfmm/aquafarm-management-service-migration go run migrations/main.go migrate

db_rollback:
	docker compose -f docker-compose.migration.yaml build
	$(if $(strip $(version)), docker run --rm naufalfmm/aquafarm-management-service-migration go run migrations/main.go rollback --version $(version), docker run --rm naufalfmm/aquafarm-management-service-migration go run migrations/main.go rollback)

db_run: 
	db_init 
	db_migrate

service:
	docker compose -f docker-compose.service.yaml up -d

test:
	go test ./... -count=1 -failfast