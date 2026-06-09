# Task API

A production-oriented REST API for task management, built with Go's standard library.

## Concepts (learning)

- HTTP server with `net/http` (ServeMux, Handler, HandlerFunc)
- JSON encoding/decoding with `encoding/json`
- Custom error types (implementing the `error` interface)
- Thread-safe in-memory store with `sync.RWMutex`
- Middleware pattern
- Structured logging with `log/slog`
- Testing HTTP handlers with `httptest`

## Project layout

```
task-api/
├── cmd/
│   └── main.go      # Entry point, server setup
├── go.mod
├── .gitignore
└── README.md
```

*(Layout will grow as the project evolves.)*

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
