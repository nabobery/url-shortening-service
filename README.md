# URL Shortening Service

https://roadmap.sh/projects/url-shortening-service

A modern URL shortening service built with Go, featuring a clean architecture and PostgreSQL for data persistence.

## Features

- Create shortened URLs
- Redirect to original URLs
- View URL statistics (access count, creation date)
- Update and delete shortened URLs
- Clean architecture with dependency injection
- PostgreSQL for data persistence
- Configurable via environment variables

## Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Docker (optional)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/nabobery/url-shortening-service.git
   cd url-shortening-service
   ```

2. Copy the environment file and configure it:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

- `POST /api/shorten` - Create a shortened URL

  ```json
  {
    "url": "https://example.com/very/long/url"
  }
  ```

- `GET /:shortCode` - Redirect to original URL
- `GET /api/urls/:shortCode` - Get URL statistics
- `PUT /api/urls/:shortCode` - Update URL
- `DELETE /api/urls/:shortCode` - Delete URL

## Project Structure

```
url-shortener/
├── cmd/                  # Application entrypoint
├── internal/
│   ├── config/              # Configuration
│   ├── core/                # Core domain logic
│   │   ├── model/          # Domain models
│   │   ├── service/        # Business logic
│   │   └── repository/     # Data access interfaces
│   └── platform/           # External implementations
│       ├── database/       # Database implementations
│       ├── web/           # HTTP handlers and router
│       └── shortener/     # URL shortening utilities
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
