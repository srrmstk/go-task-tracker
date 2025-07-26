package memo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type memoUpdater interface {
	Update(ctx context.Context, id uuid.UUID, m *model.MemoUpdateDTO) error
}

func UpdateMemoHandler(mu memoUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			helpers.JsonError(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		var m model.MemoUpdateDTO
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			helpers.JsonError(w, "invalid request body", http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(m); err != nil {
			helpers.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = mu.Update(r.Context(), id, &m)
		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			helpers.JsonError(w, "memo not found", http.StatusNotFound)
		default:
			helpers.JsonError(w, "failed to update", http.StatusInternalServerError)
		}
	}
}
