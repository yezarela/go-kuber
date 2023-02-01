include .env

$(eval export $(grep -v '^#' .env | xargs -0))

DB_DSN := "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

dep:
	@go mod download

setup-local:
	@docker-compose up -d
	@sleep 5
	@psql -h localhost -p $(DB_PORT) -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname = '$(DB_NAME)'" | grep -q 1 || psql -h localhost -p $(DB_PORT) -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME)"

run:
	air

test:
	go test ./...

build:
	@go build -o main main.go

build-docker:
	@docker build --tag material . 

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