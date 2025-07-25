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

type categoryProvider interface {
	GetByID(ctx context.Context, id string) (model.Category, error)
}

func GetCategoryHandler(log *slog.Logger, cp categoryProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("id")
		_, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		res, err := cp.GetByID(r.Context(), id)
		switch {
		case err == nil:
			json.NewEncoder(w).Encode(res)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "category not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
	}
}
