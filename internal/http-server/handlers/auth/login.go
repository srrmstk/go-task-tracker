package auth

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type login interface {
	Login(ctx context.Context, dto model.UserLoginDTO) (string, error)
}

func LoginHandler(l login) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto model.UserLoginDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := l.Login(r.Context(), dto)
		if err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}
