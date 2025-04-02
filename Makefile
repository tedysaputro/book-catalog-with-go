.PHONY: build clean run db-setup test test-setup test-cleanup test-data-preload-setup dev-setup dev

build:
	go build -o bin/book-catalog ./src

clean:
	rm -rf bin/

run: build
	./bin/book-catalog

db-setup:
	docker compose up -d postgres
	@echo "Waiting for database to be ready..."
	@sleep 3
	@docker exec book_catalog_db psql -U postgres -c "DROP DATABASE IF EXISTS book_catalog;"
	@docker exec book_catalog_db psql -U postgres -c "CREATE DATABASE book_catalog;"
	@echo "Database setup completed"

dev-setup:
	docker compose up -d postgres
	@echo "Waiting for database to be ready..."
	@sleep 3
	@docker exec book_catalog_db psql -U postgres -c "DROP DATABASE IF EXISTS book_catalog_dev;"
	@docker exec book_catalog_db psql -U postgres -c "CREATE DATABASE book_catalog_dev;"
	@docker exec -i book_catalog_db psql -U postgres -d book_catalog_dev < configs/ddl.sql
	@docker exec -i book_catalog_db psql -U postgres -d book_catalog_dev < configs/init-db.sql
	@echo "Database setup completed"

dev: build dev-setup
	@echo "set environment variable"
	export DB_HOST=localhost
	export DB_PORT=5432
	export DB_USER=postgres
	export DB_PASSWORD=postgres
	export DB_NAME=book_catalog_dev
	./bin/book-catalog

test-setup:
	docker compose up -d postgres
	@echo "Waiting for database to be ready..."
	@sleep 3
	@docker exec book_catalog_db psql -U postgres -c "DROP DATABASE IF EXISTS book_catalog_test;"
	@docker exec book_catalog_db psql -U postgres -c "CREATE DATABASE book_catalog_test;"
	@echo "Test database setup completed"

test-data-preload-setup:
	@echo "Importing initial data..."
	@docker exec -i book_catalog_db psql -U postgres -d book_catalog_test < configs/init-db.sql
	@echo "Initial data imported"

test-cleanup:
	@docker exec book_catalog_db psql -U postgres -c "DROP DATABASE IF EXISTS book_catalog_test;"

test: test-setup
	@make test-data-preload-setup
	go test -v ./tests/...
	@make test-cleanup

.DEFAULT_GOAL := build
