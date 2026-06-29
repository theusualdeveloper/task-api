package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/theusualdeveloper/task-api/formvalidation"
	"github.com/theusualdeveloper/task-api/store"
)

type CreateRequestDTO struct {
	Title string `json:"title"`
}

type TaskHandler struct {
	Logger    *slog.Logger
	TaskStore store.TaskStorer
}

func NewTaskHandler(taskStore store.TaskStorer, logger *slog.Logger) *TaskHandler {
	return &TaskHandler{
		Logger:    logger,
		TaskStore: taskStore,
	}
}

func (h *TaskHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var reqDTO CreateRequestDTO
	err := json.NewDecoder(r.Body).Decode(&reqDTO)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if reqDTO.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task := h.TaskStore.Add(reqDTO.Title)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *TaskHandler) GetListHandler(w http.ResponseWriter, r *http.Request) {
	tasks := h.TaskStore.GetAll()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *TaskHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	validation := formvalidation.NewFormValidation(r)
	id, ok := validation.ValidateID("id")
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, found := h.TaskStore.GetByID(id)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *TaskHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	validation := formvalidation.NewFormValidation(r)
	id, ok := validation.ValidateID("id")
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	deleted := h.TaskStore.Delete(id)
	if !deleted {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) CompleteHandler(w http.ResponseWriter, r *http.Request) {
	validation := formvalidation.NewFormValidation(r)
	id, ok := validation.ValidateID("id")
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, found := h.TaskStore.Update(id)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		h.Logger.Error(err.Error())
	}
}
