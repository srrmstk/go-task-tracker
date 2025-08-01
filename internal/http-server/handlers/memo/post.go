package memo

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type memoCreator interface {
	Create(ctx context.Context, m model.MemoCreateDTO) (*model.Memo, error)
}

func CreateMemoHandler(mc memoCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto model.MemoCreateDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := mc.Create(r.Context(), dto)
		if err != nil {
			helpers.JsonError(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}
