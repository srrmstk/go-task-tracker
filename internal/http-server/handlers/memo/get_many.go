package memo

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/model"
	"log/slog"
	"net/http"
)

type memosProvider interface {
	GetAll(ctx context.Context) ([]model.Memo, error)
}

func GetMemosHandler(log *slog.Logger, mp memosProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		memos, err := mp.GetAll(r.Context())
		if err != nil {
			log.Debug(err.Error())
			http.Error(w, "could not fetch memos", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(memos)
	}
}
