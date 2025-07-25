package category

import (
	"context"
	"database/sql"
	"errors"
	"go-task-tracker/internal/http-server/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type categoryDelete interface {
	Delete(ctx context.Context, id uuid.UUID) error
}

func DeleteCategoryHandler(cd categoryDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			helpers.JsonError(w, "invalid UUID format", http.StatusBadRequest)
			return
		}
		err = cd.Delete(r.Context(), id)

		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			helpers.JsonError(w, "category not found", http.StatusNotFound)
		default:
			helpers.JsonError(w, "failed to delete", http.StatusInternalServerError)
		}
	}
}
