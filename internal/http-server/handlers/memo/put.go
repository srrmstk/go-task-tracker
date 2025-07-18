package memo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type memoUpdater interface {
	Update(ctx context.Context, id int64, m *model.MemoUpdate) error
}

func UpdateMemoHandler(log *slog.Logger, mu memoUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid Id", http.StatusBadRequest)
			return
		}

		var m model.MemoUpdate
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
