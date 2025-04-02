# Book Catalog API

A RESTful API for managing a book catalog, built with Go and PostgreSQL.

## Prerequisites

- Go 1.x or higher
- PostgreSQL 16 or higher
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/tedysaputro/book-catalog-with-go.git
cd book-catalog-with-go
```

2. Run the development setup:
```bash
make dev
```

This will:
- Start PostgreSQL database using Docker Compose
- Create development database
- Import initial data
- Start the application

The API will be available at `http://localhost:8080`.

## API Endpoints

### Authors

- `GET /authors` - List all authors
  - Query Parameters:
    - `p` (page number, default: 1)
    - `limit` (items per page, default: 10)
    - `sortBy` (sort field, default: "id")
    - `direction` (sort direction: "asc" or "desc", default: "asc")
    - `authorName` (filter by author name, case-insensitive, default: "")

- `GET /authors/:id` - Get author by ID
- `POST /authors` - Create new author
  ```json
  {
    "name": "Author Name",
    "description": "Author Description"
  }
  ```

### Publishers

- `GET /publishers` - List all publishers
  - Query Parameters:
    - `p` (page number, default: 1)
    - `limit` (items per page, default: 10)
    - `sortBy` (sort field, default: "id")
    - `direction` (sort direction: "asc" or "desc", default: "asc")
    - `publisherName` (filter by publisher name, case-insensitive, default: "")

- `GET /publishers/:id` - Get publisher by ID
- `POST /publishers` - Create new publisher
  ```json
  {
    "name": "Publisher Name",
    "description": "Publisher Description"
  }
  ```

## Project Structure

```
.
├── bin/                # Compiled binaries
├── configs/           # Configuration files
│   ├── ddl.sql       # Database schema
│   └── init-db.sql   # Initial data
├── src/               # Source code
│   ├── author/        # Author domain
│   │   ├── author.go          # Author model and database operations
│   │   ├── author_dto.go      # Data Transfer Objects
│   │   ├── author_handler.go  # HTTP handlers
│   │   └── author_service.go  # Business logic
│   ├── publisher/     # Publisher domain
│   │   ├── publisher.go       # Publisher model and database operations
│   │   └── publisher_service.go # Business logic
│   ├── database.go    # Database configuration
│   └── main.go        # Application entry point
├── tests/             # Test files
├── compose.yml        # Docker Compose configuration
├── go.mod            # Go modules file
└── Makefile          # Build automation
```

## Development

### Available Make Commands

- `make dev` - Setup development environment and run the application
- `make dev-setup` - Setup development database and import initial data
- `make test` - Run tests
- `make build` - Build the application binary
- `make clean` - Clean build artifacts

## Testing

Run the tests using:
```bash
make test
```

This will:
- Create a test database
- Import test data
- Run the tests
- Clean up test database

## License

This project is open source and available under the [MIT License](LICENSE).
