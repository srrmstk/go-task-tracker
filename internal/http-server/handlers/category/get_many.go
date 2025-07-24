package category

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"
)

type categoriesProvider interface {
	GetAll(ctx context.Context) ([]model.Category, error)
}

func GetCategoriesHandler(log *slog.Logger, cp categoriesProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		res, err := cp.GetAll(r.Context())
		if err != nil {
			log.Debug(err.Error())
			http.Error(w, "could not fetch categories", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}
