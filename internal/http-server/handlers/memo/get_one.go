package memo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type memoProvider interface {
	GetByID(ctx context.Context, id string) (model.Memo, error)
}

func GetOneMemoHandler(log *slog.Logger, mp memoProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("id")
		_, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "invalid UUID format", http.StatusBadRequest)
			return
		}

		memo, err := mp.GetByID(r.Context(), id)
		switch {
		case err == nil:
			json.NewEncoder(w).Encode(memo)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "memo not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}

	}
}
