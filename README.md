# Students API

Minimal REST API for managing students using Go, `net/http`, and SQLite. It supports creating a student and fetching by id or listing all students. Configuration is provided via YAML or environment variables.

## Features
- HTTP server with graceful shutdown.
- SQLite-backed storage with auto-migration for `students` table.
- Request validation (required: `name`, `email`, `age`).
- Consistent JSON error payloads.

## Project Structure
- `cmd/students-api/main.go` — wiring: config load, storage init, routes, graceful shutdown.
- `config/` — YAML configuration (e.g., `local.yaml`).
- `internal/config` — config loading via `cleanenv` with flag/env support.
- `internal/http/handlers/student` — HTTP handlers for create/read/list.
- `internal/storage/sqlite` — SQLite implementation of storage interface.
- `internal/types` — DTOs (e.g., `Student`).
- `internal/utils/response` — JSON response helpers.
- `storage/` — default path for SQLite DB file.

## Requirements
- Go 1.25+
- SQLite (uses `modernc.org/sqlite` driver; no external binary needed)

## Configuration
Config can be supplied via `CONFIG_PATH` env var or `--config` flag pointing to a YAML file.


## Run Locally
From repo root:
```bash
# Windows PowerShell example
$env:CONFIG_PATH="config/local.yaml"
go run ./cmd/students-api

# or with flag
go run ./cmd/students-api --config config/local.yaml
```
Server starts on the configured `address` and logs startup info.

## API
Base URL: `http://<address>/api`

### Create Student
- **POST** `/students`
- Body
```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "age": 21
}
```
- Responses
  - `201 Created`: `{ "id": <int64> }`
  - `400 Bad Request`: validation error (e.g., missing field)
  - `500 Internal Server Error`

### Get Student By ID
- **GET** `/students/{id}`
- Responses
  - `200 OK`: `Student` JSON
  - `400 Bad Request`: invalid id format
  - `500 Internal Server Error`: on lookup failure

### List Students
- **GET** `/students`
- Responses
  - `200 OK`: `Student[]`
  - `500 Internal Server Error`

### Error Payloads
Errors follow:
```json
{
  "status": "Error",
  "error": "field name is required"
}
```

## Validation Rules
- `name`: required
- `email`: required (format not enforced yet)
- `age`: required

## Development Notes
- Graceful shutdown allows in-flight requests up to 5s to finish when the process receives SIGINT/SIGTERM.
- Storage implementation lives behind the `Storage` interface; swap in other backends by implementing `CreateStudent`, `GetStudentById`, `GetStudentList`.
- Logging uses `log/slog`; adjust verbosity via environment as needed.

## Testing Ideas
- Unit-test handler validation using `httptest`.
- Integration-test storage with a temp SQLite file.

