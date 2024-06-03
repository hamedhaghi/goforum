.PHONY: up up-s down lint migrate shell test

up:
	docker-compose up -d --build

up-s:
	@docker-compose up -d --build > /dev/null 2>&1

down:
	docker-compose down

lint:
	@docker-compose exec -it app golangci-lint run

migrate:
	@docker-compose exec -it app go run migrations/migrate.go 

shell:
	@docker-compose exec -it app sh

test:
	@go test ./...



		
			