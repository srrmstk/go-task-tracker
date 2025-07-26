package memo

import (
	"context"
	"database/sql"
	"errors"
	"go-task-tracker/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type memoDelete interface {
	Delete(ctx context.Context, id uuid.UUID) error
}

func DeleteMemoHandler(md memoDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			helpers.JsonError(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		err = md.Delete(r.Context(), id)

		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			helpers.JsonError(w, "Memo not found", http.StatusNotFound)
		default:
			helpers.JsonError(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}
