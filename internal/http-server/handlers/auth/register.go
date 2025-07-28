package auth

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type register interface {
	Register(ctx context.Context, dto model.UserRegisterDTO) (*model.UserRegisterResponseDTO, error)
}

func RegisterHandler(reg register) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto model.UserRegisterDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := reg.Register(r.Context(), dto)
		if err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}
