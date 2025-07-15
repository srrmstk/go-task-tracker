package task

import (
	"fmt"
	"net/http"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task ID not provided!"))
		return
	}

	response := fmt.Sprintf("Task ID: %v deleted", id)
	w.Write([]byte(response))
}
