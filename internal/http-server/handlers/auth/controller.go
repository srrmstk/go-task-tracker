package auth

import (
	"go-task-tracker/internal/service"

	"github.com/go-chi/chi/v5"
)

type authController struct {
	s *service.AuthService
}

func NewAuthController(s *service.AuthService) *authController {
	return &authController{s: s}
}

func (c *authController) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", RegisterHandler(c.s))
		r.Post("/login", LoginHandler(c.s))
		r.Post("/verify/{id}", VerifyHandler(c.s))
	})
}
