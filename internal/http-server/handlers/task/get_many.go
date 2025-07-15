package task

import (
	"encoding/json"
	"go-task-tracker/internal/service"
	"net/http"
)

func GetTasksHandler(s *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := s.GetAll(r.Context())
		if err != nil {
			http.Error(w, "could not fetch tasks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}
