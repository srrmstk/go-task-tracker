package memo

import (
	"context"
	"encoding/json"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"net/http"
)

type memosProvider interface {
	GetAll(ctx context.Context) ([]model.Memo, error)
}

func GetMemosHandler(mp memosProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memos, err := mp.GetAll(r.Context())
		if err != nil {
			helpers.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(memos)
	}
}
