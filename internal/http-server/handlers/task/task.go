package task

import (
	"go-task-tracker/internal/http-server/middleware/logger"
	"go-task-tracker/internal/service"
	"log/slog"
	"net/http"
)

type taskController struct {
	s *service.TaskService
}

func NewTaskController(s *service.TaskService) *taskController {
	return &taskController{s: s}
}

func (c *taskController) Register(mux *http.ServeMux, log *slog.Logger) {
	mux.HandleFunc("GET /tasks", logger.LoggerMiddleware(log, GetTasksHandler(c.s)))
	mux.HandleFunc("GET /tasks/{id}", logger.LoggerMiddleware(log, GetOneTaskHandler(log, c.s)))
	mux.HandleFunc("POST /tasks", logger.LoggerMiddleware(log, CreateTaskHandler(log, c.s)))
	mux.HandleFunc("PUT /tasks/{id}", logger.LoggerMiddleware(log, UpdateTaskHandler(log, c.s)))
	// mux.HandleFunc("DELETE /tasks/{id}", logger.LoggerMiddleware(log, DeleteTaskHandler))
}
