package task

import (
	"go-task-tracker/internal/http-server/middleware/logger"
	"go-task-tracker/internal/service"
	"log/slog"
	"net/http"
)

type TaskController struct {
	s *service.TaskService
}

func NewTaskController(s *service.TaskService) *TaskController {
	return &TaskController{s: s}
}

func (t *TaskController) Register(mux *http.ServeMux, log *slog.Logger) {
	mux.HandleFunc("GET /tasks", logger.LoggerMiddleware(log, GetTasksHandler(t.s)))
	// mux.HandleFunc("GET /tasks/{id}", logger.LoggerMiddleware(log, GetOneTaskHandler(log, t.s)))
	mux.HandleFunc("POST /tasks", logger.LoggerMiddleware(log, CreateTaskHandler(log, t.s)))
	// mux.HandleFunc("PUT /tasks/{id}", logger.LoggerMiddleware(log, PutTaskHandler))
	// mux.HandleFunc("DELETE /tasks/{id}", logger.LoggerMiddleware(log, DeleteTaskHandler))
}
