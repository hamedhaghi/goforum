.PHONY: *

db=./database/forum.db
app=app

up:
	mkdir -p $(dir $(db)) && \
	touch $(db) && \
	docker compose build --no-cache && \
	docker compose up -d && \
	$(MAKE) migrate
	docker compose logs -f

down:
	rm -rf $(db) && \
	docker compose down

migrate:
	docker compose exec $(app) go run migrations/migrate.go 

shell:
	docker compose exec -it $(app) sh
