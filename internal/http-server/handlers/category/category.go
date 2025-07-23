package category

import (
	"go-task-tracker/internal/http-server/middleware/logger"
	"go-task-tracker/internal/service"
	"log/slog"
	"net/http"
)

type categoryController struct {
	s *service.CategoryService
}

func NewCategoryController(s *service.CategoryService) *categoryController {
	return &categoryController{s: s}
}

func (c *categoryController) Register(mux *http.ServeMux, log *slog.Logger) {
	mux.HandleFunc("GET /categories/", logger.LoggerMiddleware(log, GetCategoriesHandler(log, c.s)))
	mux.HandleFunc("GET /categories/{id}", logger.LoggerMiddleware(log, GetCategoryHandler(log, c.s)))
	mux.HandleFunc("POST /categories/", logger.LoggerMiddleware(log, CreateCategoryHandler(log, c.s)))
	mux.HandleFunc("PUT /categories/{id}", logger.LoggerMiddleware(log, UpdateCategoryHandler(log, c.s)))
	mux.HandleFunc("DELETE /categories/{id}", logger.LoggerMiddleware(log, DeleteCategoryHandler(log, c.s)))
}
