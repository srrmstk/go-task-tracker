package category

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

type categoryProvider interface {
	GetByID(ctx context.Context, id int64) (model.Category, error)
}

func GetCategoryHandler(log *slog.Logger, cp categoryProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		strId := r.PathValue("id")
		id, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			http.Error(w, "invalid ID", http.StatusBadRequest)
			return
		}

		res, err := cp.GetByID(r.Context(), id)
		switch {
		case err == nil:
			json.NewEncoder(w).Encode(res)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "category not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
	}
}
