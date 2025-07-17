package task

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"
)

type tasksProvider interface {
	GetAll(ctx context.Context) ([]model.Task, error)
}

func GetTasksHandler(log *slog.Logger, tp tasksProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		tasks, err := tp.GetAll(r.Context())
		if err != nil {
			log.Debug(err.Error())
			http.Error(w, "could not fetch tasks", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)
	}
}
