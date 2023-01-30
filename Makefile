include .env

$(eval export $(grep -v '^#' .env | xargs -0))

DB_DSN := "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

test:
	@echo ${DB_USER}
run:
	air

migrate-create:
	@read -p "Please provide name for the migration: " Name; \
	migrate create -ext sql -dir infra/database/migrations $${Name}
	
migrate-up:
	@read -p "How many migration you wants to perform (default: [all]): " N; \
	migrate -database $(DB_DSN) -path=infra/database/migrations up ${NN}

migrate-down:
	@read -p "How many migration you wants to perform (default: [all]): " N; \
	migrate -database $(DB_DSN) -path=infra/database/migrations down ${NN}

migrate-drop:
	migrate -database $(DB_DSN) -path=infra/database/migrations drop


# mockgen -source=product/repository/product_repository.go -destination product/mocks/product_repository_mock.go -package=mock