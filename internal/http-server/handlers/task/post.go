package task

import (
	"encoding/json"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/service"
	"log/slog"
	"net/http"
)

func CreateTaskHandler(log *slog.Logger, s *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t model.Task
		log.Debug("GetTaskHandler", "body", t)

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.Create(r.Context(), &t); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
	}
}
