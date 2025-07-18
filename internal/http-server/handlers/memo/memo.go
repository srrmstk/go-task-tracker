package memo

import (
	"go-task-tracker/internal/http-server/middleware/logger"
	"go-task-tracker/internal/service"
	"log/slog"
	"net/http"
)

type memoController struct {
	s *service.MemoService
}

func NewMemoController(s *service.MemoService) *memoController {
	return &memoController{s: s}
}

func (c *memoController) Register(mux *http.ServeMux, log *slog.Logger) {
	mux.HandleFunc("GET /tasks", logger.LoggerMiddleware(log, GetMemosHandler(log, c.s)))
	mux.HandleFunc("GET /tasks/{id}", logger.LoggerMiddleware(log, GetOneMemoHandler(log, c.s)))
	mux.HandleFunc("POST /tasks", logger.LoggerMiddleware(log, CreateMemoHandler(log, c.s)))
	mux.HandleFunc("PUT /tasks/{id}", logger.LoggerMiddleware(log, UpdateMemoHandler(log, c.s)))
	mux.HandleFunc("DELETE /tasks/{id}", logger.LoggerMiddleware(log, DeleteMemoHandler(log, c.s)))
}
