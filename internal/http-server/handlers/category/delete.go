package category

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
)

type categoryDelete interface {
	Delete(ctx context.Context, id int64) error
}

func DeleteCategoryHandler(log *slog.Logger, cd categoryDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		err = cd.Delete(r.Context(), id)

		switch {
		case err == nil:
			w.WriteHeader(http.StatusNoContent)
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "category not found", http.StatusNotFound)
		default:
			log.Debug(err.Error())
			http.Error(w, "failed to delete", http.StatusInternalServerError)
		}
	}
}
