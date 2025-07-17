package task

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"
)

type taskCreator interface {
	Create(ctx context.Context, m *model.Task) error
}

func CreateTaskHandler(log *slog.Logger, tc taskCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var t model.Task
		log.Debug("GetTaskHandler", "body", t)

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := tc.Create(r.Context(), &t); err != nil {
			log.Debug(err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
	}
}
