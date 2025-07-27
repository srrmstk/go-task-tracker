package memo

import (
	"go-task-tracker/internal/http-server/middleware"
	"go-task-tracker/internal/service"

	"github.com/go-chi/chi/v5"
)

type memoController struct {
	s *service.MemoService
}

func NewMemoController(s *service.MemoService) *memoController {
	return &memoController{s: s}
}

func (c *memoController) Register(r chi.Router) {
	r.Route("/memos", func(r chi.Router) {
		r.Use(middleware.JwtGuard)
		r.Get("/", GetMemosHandler(c.s))
		r.Get("/{id}", GetOneMemoHandler(c.s))
		r.Post("/", CreateMemoHandler(c.s))
		r.Put("/{id}", UpdateMemoHandler(c.s))
		r.Delete("/{id}", DeleteMemoHandler(c.s))
	})
}
