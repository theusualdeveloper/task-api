package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/theusualdeveloper/task-api/handler"
	"github.com/theusualdeveloper/task-api/store"
)

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name       string
		req        handler.CreateRequestDTO
		wantStatus int
	}{
		{
			name: "Test 1",
			req: handler.CreateRequestDTO{
				Title: "test title",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "Test 2",
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.req)
			if err != nil {
				t.Fatalf("marshaling request data failed: %s", err.Error())
			}
			r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(b))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			logger := slog.New(slog.NewTextHandler(io.Discard, nil))
			h := handler.NewTaskHandler(
				store.NewTaskStore(),
				logger,
			)
			h.CreateHandler(w, r)
			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Fatalf("want status code: %d, got: %d", tt.wantStatus, res.StatusCode)
			}
		})
	}
}

func TestGetListHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	h := handler.NewTaskHandler(
		store.NewTaskStore(),
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	h.GetListHandler(w, r)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("want status code: %d, got: %d", http.StatusOK, res.StatusCode)
	}
}

func TestGetByIDHandler(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		wantId     int
		wantStatus int
	}{
		{
			name:       "test 1",
			title:      "title 1",
			wantId:     1,
			wantStatus: http.StatusOK,
		},
		{
			name:       "test 2",
			title:      "title 2",
			wantId:     10,
			wantStatus: http.StatusNotFound,
		},
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := store.NewTaskStore()
			h := handler.NewTaskHandler(
				ts,
				logger,
			)
			ts.Add(tt.title)
			r := httptest.NewRequest(http.MethodGet, "/tasks/", nil)
			r.SetPathValue("id", strconv.Itoa(tt.wantId))
			w := httptest.NewRecorder()
			h.GetByIDHandler(w, r)
			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Fatalf("want status code: %d, got: %d", tt.wantStatus, res.StatusCode)
			}
		})
	}
}

func TestDeleteHandler(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		wantId     int
		wantStatus int
	}{
		{
			name:       "test 1",
			title:      "title 1",
			wantId:     1,
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "test 2",
			title:      "title 2",
			wantId:     10,
			wantStatus: http.StatusNotFound,
		},
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := store.NewTaskStore()
			h := handler.NewTaskHandler(ts, logger)
			r := httptest.NewRequest(http.MethodDelete, "/tasks/", nil)
			r.SetPathValue("id", strconv.Itoa(tt.wantId))
			w := httptest.NewRecorder()
			ts.Add(tt.title)
			h.DeleteHandler(w, r)
			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Fatalf("want status code: %d, got: %d", tt.wantStatus, res.StatusCode)
			}
		})
	}
}

func TestCompleteHandler(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		id         int
		wantStatus int
	}{
		{
			name:       "Test 1",
			title:      "test title 1",
			id:         1,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Test 2",
			title:      "test title 2",
			id:         10,
			wantStatus: http.StatusNotFound,
		},
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := store.NewTaskStore()
			h := handler.NewTaskHandler(
				ts,
				logger,
			)
			r := httptest.NewRequest(http.MethodPatch, "/", nil)
			r.SetPathValue("id", strconv.Itoa(tt.id))
			w := httptest.NewRecorder()
			ts.Add(tt.title)
			h.CompleteHandler(w, r)
			res := w.Result()
			var task store.Task
			err := json.NewDecoder(res.Body).Decode(&task)
			if err != nil {
				t.Fatal("json decode failed")
			}
			if res.StatusCode != tt.wantStatus {
				t.Fatalf("want status code: %d, got: %d", tt.wantStatus, res.StatusCode)
			}
			if tt.wantStatus == http.StatusOK && !task.Done {
				t.Fatal("task must be done when status code is 200")
			}
		})
	}
}
