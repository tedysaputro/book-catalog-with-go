.PHONY: build clean run test test-setup test-cleanup

build:
	go build -o bin/book-catalog ./src

clean:
	rm -rf bin/

run: build
	./bin/book-catalog

test-setup:
	docker compose up -d postgres
	@echo "Waiting for database to be ready..."
	@sleep 3
	@docker exec book_catalog_db psql -U postgres -c "DROP DATABASE IF EXISTS book_catalog_test;"
	@docker exec book_catalog_db psql -U postgres -c "CREATE DATABASE book_catalog_test;"

test-cleanup:
	@docker exec book_catalog_db psql -U postgres -c "DROP DATABASE IF EXISTS book_catalog_test;"

test: test-setup
	go test -v ./tests/...
	@make test-cleanup

.DEFAULT_GOAL := build
