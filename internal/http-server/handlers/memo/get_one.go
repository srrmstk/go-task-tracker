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
)

type MemoProvider interface {
	GetByID(ctx context.Context, id int64) (model.Memo, error)
}

func GetOneMemoHandler(log *slog.Logger, mp MemoProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		strID := r.PathValue("id")
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			http.Error(w, "invalid ID", http.StatusBadRequest)
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
