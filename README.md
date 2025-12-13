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

## API Documentation

This section provides detailed API documentation for the Students API.

### Base URL
The base URL for the API is: `http://localhost:8082/api`

### Endpoints

#### Create Student
- **POST** `/students`
- **Description**: Creates a new student.
- **Request Body**:
  ```json
  {
    "name": "Jane Doe",
    "email": "jane@example.com",
    "age": 21
  }
  ```
- **Responses**:
  - `201 Created`: Returns the created student ID.
    ```json
    {
      "id": <int64>
    }
    ```
  - `400 Bad Request`: Validation error (e.g., missing field).
  - `500 Internal Server Error`: On server error.

#### Get Student By ID
- **GET** `/students/{id}`
- **Description**: Retrieves a student by their ID.
- **Responses**:
  - `200 OK`: Returns the student details.
    ```json
    {
      "id": <int64>,
      "name": "Jane Doe",
      "email": "jane@example.com",
      "age": 21
    }
    ```
  - `400 Bad Request`: Invalid ID format.
  - `500 Internal Server Error`: On lookup failure.

#### List Students
- **GET** `/students`
- **Description**: Retrieves a list of all students.
- **Responses**:
  - `200 OK`: Returns an array of students.
    ```json
    [
      {
        "id": <int64>,
        "name": "Jane Doe",
        "email": "jane@example.com",
        "age": 21
      }
    ]
    ```
  - `500 Internal Server Error`: On server error.

### Error Handling
All error responses follow this structure:
```json
{
  "status": "Error",
  "error": "<error message>"
}
```

### Validation Rules
- `name`: required
- `email`: required (format not enforced yet)
- `age`: required

### Development Notes
- The API supports graceful shutdown, allowing in-flight requests to finish when the server is stopped.
- The storage implementation is abstracted behind the `Storage` interface, allowing for easy swapping of storage backends.

### Testing
- Unit tests should be written for handler validation using `httptest`.
- Integration tests should be conducted with a temporary SQLite file.

