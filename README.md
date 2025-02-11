# Book Store API

A simple RESTful API built with Go and Gin framework for managing a book store.

## Features

- CRUD operations for books
- Swagger documentation
- Health check endpoint
- UUID-based identifiers
- In-memory storage

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-rest-api
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## API Documentation

Swagger documentation is available at:
- `http://localhost:8080` (redirects to Swagger UI)
- `http://localhost:8080/swagger/index.html`

## API Endpoints

- `GET /health` - Health check
- `GET /api/v1/books` - List all books
- `GET /api/v1/books/:id` - Get a book by ID
- `POST /api/v1/books` - Create a new book
- `PUT /api/v1/books/:id` - Update a book
- `DELETE /api/v1/books/:id` - Delete a book

## Example Request

Create a new book:
```bash
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan & Brian W. Kernighan",
    "price": 49.99
  }'
```

## License

MIT
