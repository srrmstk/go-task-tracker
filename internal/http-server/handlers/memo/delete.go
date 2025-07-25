package memo

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type memoDelete interface {
	Delete(ctx context.Context, id string) error
}

func DeleteMemoHandler(log *slog.Logger, md memoDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		id := r.PathValue("id")
		_, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		err = md.Delete(r.Context(), id)

		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "Memo not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}
