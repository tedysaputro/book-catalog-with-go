.PHONY: build clean run

build:
	go build -o bin/book-catalog ./src

clean:
	rm -rf bin/

run: build
	./bin/book-catalog

.DEFAULT_GOAL := build
