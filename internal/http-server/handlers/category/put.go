package category

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/http-server/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type categoryUpdater interface {
	Update(ctx context.Context, id uuid.UUID, c *model.CategoryUpdateDTO) error
}

func UpdateCategoryHandler(cu categoryUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			helpers.JsonError(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		var c model.CategoryUpdateDTO
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = cu.Update(r.Context(), id, &c)
		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			helpers.JsonError(w, "category not found", http.StatusNotFound)
		default:
			helpers.JsonError(w, "failed to update", http.StatusInternalServerError)
		}
	}
}
