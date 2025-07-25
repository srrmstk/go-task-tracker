package category

import (
	"go-task-tracker/internal/service"

	"github.com/go-chi/chi/v5"
)

type categoryController struct {
	s *service.CategoryService
}

func NewCategoryController(s *service.CategoryService) *categoryController {
	return &categoryController{s: s}
}

func (c *categoryController) Register(r chi.Router) {
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", GetCategoriesHandler(c.s))
		r.Get("/{id}", GetCategoryHandler(c.s))
		r.Post("/", CreateCategoryHandler(c.s))
		r.Put("/{id}", UpdateCategoryHandler(c.s))
		r.Delete("/{id}", DeleteCategoryHandler(c.s))
	})
}
