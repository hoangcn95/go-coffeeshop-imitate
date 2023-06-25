# include .env
# export

#postgres local
POSTGRES_USER=postgres
POSTGRES_PASSWORD=P@ssw0rd
POSTGRES_DB=postgres
POSTGRES_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable

#sql-migration gen
SQL_MIGRATION_DIR=db/migrations

# migration
migrategen:
	@echo "====== Migrate gen name: $(NAME) ======"
	@migrate create -ext sql -seq -dir $(SQL_MIGRATION_DIR) $(NAME)
.PHONY: migrategen

migrateup:
	@echo "====== START MIGRATION UP ======"
	@migrate -path $(SQL_MIGRATION_DIR) -database "$(POSTGRES_URL)" -verbose up
.PHONY: migrateup

migratedown:
	@echo "====== START MIGRATION DOWN ======"
	@migrate -path $(SQL_MIGRATION_DIR) -database "$(POSTGRES_URL)" -verbose down
.PHONY: migratedown

# sqlc
sqlc:
	sqlc generate
.PHONY: sqlc

# buf generator
bufall: buflint bufgen

buflint:
	buf lint
.PHONY: buflint

bufgen:
	buf generate
.PHONY: bufgen

# wire generate dependency injecttion
wire:
	cd internal/counter/app && wire && cd -
.PHONY: wire

# docker
docker-compose-core: docker-compose-core-stop docker-compose-core-start

docker-compose-core-start:
	docker-compose -f docker-compose-core.yaml up --build -d
.PHONY: docker-compose-core-start

docker-compose-core-stop:
	docker-compose -f docker-compose-core.yaml down --remove-orphans -v
.PHONY: docker-compose-core-stop