package category

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type categoryDelete interface {
	Delete(ctx context.Context, id string) error
}

func DeleteCategoryHandler(log *slog.Logger, cd categoryDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("id")
		_, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "invalid UUID format", http.StatusBadRequest)
			return
		}
		err = cd.Delete(r.Context(), id)

		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "category not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "failed to delete", http.StatusInternalServerError)
		}
	}
}
