include .env

db="${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

up:
	docker-compose up -d --build

up-s:
	@docker-compose up -d --build > /dev/null 2>&1

down:
	docker-compose down

lint:
	@docker-compose exec -it app golangci-lint run

migrate:
	@go run migrations/migrate.go 



		
			