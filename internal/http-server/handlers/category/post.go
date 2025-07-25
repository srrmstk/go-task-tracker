package category

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type categoryCreator interface {
	Create(ctx context.Context, c model.CategoryCreateDTO) (*model.Category, error)
}

func CreateCategoryHandler(log *slog.Logger, cc categoryCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dto model.CategoryCreateDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := cc.Create(r.Context(), dto)
		if err != nil {
			log.Debug(err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}
