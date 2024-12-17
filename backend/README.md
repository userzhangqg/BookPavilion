# BookPavilion Backend

The backend service for BookPavilion, an online book storage and reading platform.

## Features

- Book upload and storage
- Book metadata management
- File format validation
- Pagination support
- RESTful API

## API Endpoints

### Books

```
POST   /api/books      - Upload a new book
GET    /api/books      - List books (with pagination)
GET    /api/books/:id  - Get book details
DELETE /api/books/:id  - Delete a book
```

## Development Setup

### Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- Make (optional, for using Makefile commands)

### Environment Variables

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=bookpavilion

# Server Configuration
PORT=8080
GIN_MODE=debug  # Use 'release' for production
UPLOAD_DIR=./uploads
```

### Getting Started

1. Clone the repository
```bash
git clone <repository-url>
cd backend
```

2. Install dependencies
```bash
make deps
# or
go mod download
```

3. Create the database
```bash
make db-create
# or manually create 'bookpavilion' and 'bookpavilion_test' databases
```

4. Run the application
```bash
make run
# or
go run main.go
```

### Testing

Run all tests:
```bash
make test
# or
go test ./...
```

## Project Structure

```
backend/
├── config/         - Configuration management
├── controllers/    - HTTP request handlers
├── models/         - Data models
├── services/       - Business logic
├── uploads/        - Uploaded files directory
├── main.go         - Application entry point
├── go.mod          - Go module file
└── Makefile        - Build and development commands
```

## API Usage Examples

### Upload a Book

```bash
curl -X POST http://localhost:8080/api/books \
  -F "title=Sample Book" \
  -F "author=John Doe" \
  -F "file=@/path/to/book.pdf"
```

### List Books

```bash
curl http://localhost:8080/api/books?page=1&page_size=10
```

### Get Book Details

```bash
curl http://localhost:8080/api/books/1
```

### Delete a Book

```bash
curl -X DELETE http://localhost:8080/api/books/1
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- 200: Success
- 201: Created
- 400: Bad Request (invalid input)
- 404: Not Found
- 500: Internal Server Error

Error response format:
```json
{
  "error": "Error message here"
}
```

## Development Guidelines

1. **Code Style**
   - Follow Go standard formatting (use `gofmt`)
   - Add comments for exported functions and types
   - Use meaningful variable and function names

2. **Testing**
   - Write unit tests for all new functionality
   - Maintain test coverage
   - Use table-driven tests where appropriate

3. **Error Handling**
   - Use meaningful error messages
   - Log errors appropriately
   - Return proper HTTP status codes

4. **Git Workflow**
   - Create feature branches
   - Write clear commit messages
   - Keep commits focused and atomic

## Build and Deployment

### Local Build

```bash
make build
```

### Docker Build

```bash
make docker-build
```

### Production Deployment Notes

1. Set appropriate environment variables
2. Use `GIN_MODE=release`
3. Configure proper database credentials
4. Set up proper file storage solution
5. Configure logging
6. Set up monitoring

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.
