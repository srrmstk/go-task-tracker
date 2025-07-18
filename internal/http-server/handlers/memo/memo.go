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
	mux.HandleFunc("GET /memos/", logger.LoggerMiddleware(log, GetMemosHandler(log, c.s)))
	mux.HandleFunc("GET /memos/{id}", logger.LoggerMiddleware(log, GetOneMemoHandler(log, c.s)))
	mux.HandleFunc("POST /memos/", logger.LoggerMiddleware(log, CreateMemoHandler(log, c.s)))
	mux.HandleFunc("PUT /memos/{id}", logger.LoggerMiddleware(log, UpdateMemoHandler(log, c.s)))
	mux.HandleFunc("DELETE /memos/{id}", logger.LoggerMiddleware(log, DeleteMemoHandler(log, c.s)))
}
