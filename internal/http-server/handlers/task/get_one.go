package task

import (
	"fmt"
	"log/slog"
	"net/http"
)

func GetOneTaskHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			log.Error("Task ID not provided")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Task ID not provided!"))
			return
		}

		response := fmt.Sprintf("Task ID: %v retrieved", id)
		w.Write([]byte(response))
	}
}
