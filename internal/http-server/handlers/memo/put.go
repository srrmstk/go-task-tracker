package memo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type memoUpdater interface {
	Update(ctx context.Context, id string, m *model.MemoUpdateDTO) error
}

func UpdateMemoHandler(log *slog.Logger, mu memoUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("id")
		_, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		var m model.MemoUpdateDTO
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(m); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = mu.Update(r.Context(), id, &m)
		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "memo not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "failed to update", http.StatusInternalServerError)
		}
	}
}
