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

type categoryUpdater interface {
	Update(ctx context.Context, id int64, c *model.CategoryUpdate) error
}

func UpdateCategoryHandler(log *slog.Logger, cu categoryUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		var c model.CategoryUpdate
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = cu.Update(r.Context(), id, &c)
		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "category not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "failed to update", http.StatusInternalServerError)
		}
	}
}
