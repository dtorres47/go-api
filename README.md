# go-api

A simple Go REST API with Postgres persistence.

## 🚀 Features

- Layered architecture (controller, service, router, db)
- JSON endpoints:
    - `GET /` → “Hello from the service layer!”
    - `GET /user` → returns first user from Postgres
    - `POST /user` → creates a new user, returns its ID
- Dockerized Postgres (via `docker-compose.yml`)
- Environment-driven configuration (`.env`)
- Linting with `golangci-lint`

## 🔧 Prerequisites

- Go 1.24+
- Docker & Docker Compose

## ⚙️ Setup

1. Clone the repo and `cd go-api`.
2. Copy `.env.example` to `.env` and fill in credentials:

   ```bash
   cp .env.example .env
   ```

   **`.env.example`:**

   ```dotenv
   DB_USER=postgres
   DB_PASS=password
   DB_NAME=go_api_db
   ```

3. Start Postgres:

   ```bash
   docker compose up -d
   ```

4. Install Go dependencies:

   ```bash
   go mod tidy
   go install github.com/joho/godotenv@latest
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

## ▶️ Running the API

1. Run the server:

   ```bash
   go run main.go
   ```

2. In your browser or via `curl`, visit:

   ```http
   http://127.0.0.1:8081/
   http://127.0.0.1:8081/user
   ```

## 📝 Lint & Format

```bash
go fmt ./...
go vet ./...
golangci-lint run
```

## 📦 Dockerized (optional)

If you want your Go app in Docker alongside Postgres, add this service to `docker-compose.yml`:

```yaml
  api:
    build: .
    volumes:
      - .:/app
    working_dir: /app
    command: go run main.go
    ports:
      - "8081:8081"
    depends_on:
      - db
```

Then:

```bash
docker compose up -d
```
