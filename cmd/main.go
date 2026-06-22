package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/theusualdeveloper/task-api/handler"
	"github.com/theusualdeveloper/task-api/middleware"
	"github.com/theusualdeveloper/task-api/store"
)

func main() {
	logger := InitSlog()
	taskStore := store.NewTaskStore()
	h := handler.NewTaskHandler(taskStore, logger)

	tasks := http.NewServeMux()
	tasks.HandleFunc("POST /", h.CreateHandler)
	tasks.HandleFunc("GET /", h.GetListHandler)
	tasks.HandleFunc("GET /{id}", h.GetByIDHandler)
	tasks.HandleFunc("DELETE /{id}", h.DeleteHandler)

	mux := http.NewServeMux()
	mux.Handle("/tasks/", middleware.SetJsonHeader(http.StripPrefix("/tasks", tasks)))
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"status": "ok",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	})
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	c := make(chan struct{})
	go func(server *http.Server, logger *slog.Logger, c chan struct{}) {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			logger.LogAttrs(
				ctx,
				slog.LevelError,
				"Shutdown failed",
				slog.String("err", "Server Shutdown failed"),
			)
		}
		logger.LogAttrs(
			ctx,
			slog.LevelInfo,
			"Shutdown Success",
			slog.String("info", "Server is shutting down gracefully..."),
		)
		close(c)
	}(&server, logger, c)
	logger.Info("Server is starting on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			logger.Info("Server shutting down successfully")
		} else {
			logger.Error(err.Error())
			return
		}
	}
	<-c
}

func InitSlog() *slog.Logger {
	options := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stderr, options)
	return slog.New(handler)
}
