package category

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/http-server/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type categoryCreator interface {
	Create(ctx context.Context, c model.CategoryCreateDTO) (*model.Category, error)
}

func CreateCategoryHandler(cc categoryCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto model.CategoryCreateDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := cc.Create(r.Context(), dto)
		if err != nil {
			helpers.JsonError(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}
