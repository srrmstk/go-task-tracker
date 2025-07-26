package category

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type categoryProvider interface {
	GetByID(ctx context.Context, id uuid.UUID) (model.Category, error)
}

func GetCategoryHandler(cp categoryProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			helpers.JsonError(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		res, err := cp.GetByID(r.Context(), id)
		switch {
		case err == nil:
			json.NewEncoder(w).Encode(res)
		case errors.Is(err, sql.ErrNoRows):
			helpers.JsonError(w, "category not found", http.StatusNotFound)
		default:
			helpers.JsonError(w, "something went wrong", http.StatusInternalServerError)
		}
	}
}
