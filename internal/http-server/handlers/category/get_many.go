package category

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/http-server/helpers"
	"go-task-tracker/internal/model"
	"net/http"
)

type categoriesProvider interface {
	GetAll(ctx context.Context) ([]model.Category, error)
}

func GetCategoriesHandler(cp categoriesProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := cp.GetAll(r.Context())
		if err != nil {
			helpers.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}
