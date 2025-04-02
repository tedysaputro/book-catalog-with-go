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

2. Start the PostgreSQL database using Docker Compose:
```bash
docker-compose up -d
```

3. Run the application:
```bash
make run
```

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
    "bio": "Author Biography"
  }
  ```

## Project Structure

```
.
├── bin/                # Compiled binaries
├── config/            # Configuration files
├── src/               # Source code
│   ├── author/        # Author domain
│   │   ├── author.go          # Author model and database operations
│   │   ├── author_dto.go      # Data Transfer Objects
│   │   ├── author_handler.go  # HTTP handlers
│   │   └── author_service.go  # Business logic
│   ├── routes.go      # API routes
│   └── main.go        # Application entry point
├── compose.yml        # Docker Compose configuration
├── go.mod            # Go modules file
└── Makefile          # Build automation
```

## Development

### Available Make Commands

- `make run` - Build and run the application
- `make build` - Build the application binary

## License

This project is open source and available under the [MIT License](LICENSE).
