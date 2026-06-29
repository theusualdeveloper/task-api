# Task API

A production-oriented REST API for task management, built with Go's standard library.

## Concepts (learning)

- HTTP server with `net/http` (ServeMux, Handler, HandlerFunc)
- JSON encoding/decoding with `encoding/json`
- Thread-safe in-memory store with `sync.RWMutex`
- Middleware pattern
- Structured logging with `log/slog`
- Graceful shutdown with `os/signal` and `context`
- Interfaces and dependency injection
- Testing HTTP handlers with `httptest` and mocks

## Endpoints

| Method   | Path            | Description        |
|----------|-----------------|--------------------|
| `GET`    | `/health`       | Health check       |
| `POST`   | `/tasks/`       | Create a task      |
| `GET`    | `/tasks/`       | List all tasks     |
| `GET`    | `/tasks/{id}`   | Get task by ID     |
| `PATCH`  | `/tasks/{id}`   | Mark task as done  |
| `DELETE` | `/tasks/{id}`   | Delete a task      |

## Project layout

```
task-api/
├── cmd/
│   └── main.go            # Entry point, server and graceful shutdown
├── handler/
│   ├── task_handler.go    # HTTP handlers
│   └── mock_task_store.go # Mock store for tests
├── store/
│   ├── task_store.go      # In-memory TaskStore with RWMutex
│   ├── task_storer.go     # TaskStorer interface
│   └── Task.go            # Task model
├── middleware/
│   └── middleware.go      # JSON Content-Type middleware
├── formvalidation/
│   └── form_validation.go # Path parameter validation
├── go.mod
├── .gitignore
└── README.md
```

## Usage

```bash
git clone https://github.com/theusualdeveloper/task-api.git
cd task-api
go run ./cmd
```

Server listens on `http://localhost:8080`.

## Requirements

- Go 1.22+

## License

MIT
