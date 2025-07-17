package task

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
)

type TaskDelete interface {
	Delete(ctx context.Context, id int64) error
}

func DeleteTaskHandler(log *slog.Logger, td TaskDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		err = td.Delete(r.Context(), id)

		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "Task not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}
