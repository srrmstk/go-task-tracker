package task

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"
	"strconv"
)

type taskUpdater interface {
	Update(ctx context.Context, id int64, m *model.TaskUpdate) error
}

func UpdateTaskHandler(log *slog.Logger, tu taskUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid Id", http.StatusBadRequest)
			return
		}

		var t model.TaskUpdate
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		err = tu.Update(r.Context(), id, &t)
		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "failed to update", http.StatusInternalServerError)
		}
	}
}
