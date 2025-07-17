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

type TaskProvider interface {
	GetByID(ctx context.Context, id int64) (model.Task, error)
}

func GetOneTaskHandler(log *slog.Logger, tg TaskProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		strID := r.PathValue("id")
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			http.Error(w, "invalid ID", http.StatusBadRequest)
			return
		}

		task, err := tg.GetByID(r.Context(), id)
		switch {
		case err == nil:
			json.NewEncoder(w).Encode(task)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}

	}
}
