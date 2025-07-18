package memo

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"
)

type memoCreator interface {
	Create(ctx context.Context, m *model.Memo) error
}

func CreateMemoHandler(log *slog.Logger, mc memoCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var m model.Memo
		log.Debug("GetMemoHandler", "body", m)

		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := mc.Create(r.Context(), &m); err != nil {
			log.Debug(err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(m)
	}
}
