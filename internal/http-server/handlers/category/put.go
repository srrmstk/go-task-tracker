package category

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type categoryUpdater interface {
	Update(ctx context.Context, id uuid.UUID, c *model.CategoryUpdateDTO) error
}

func UpdateCategoryHandler(log *slog.Logger, cu categoryUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		var c model.CategoryUpdateDTO
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = cu.Update(r.Context(), id, &c)
		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "category not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "failed to update", http.StatusInternalServerError)
		}
	}
}
