package auth

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type verifier interface {
	Verify(ctx context.Context, id uuid.UUID, dto *model.UserVerifyDTO) error
}

func VerifyHandler(s verifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			helpers.JsonError(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		var dto *model.UserVerifyDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.Verify(r.Context(), id, dto); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
	}
}
