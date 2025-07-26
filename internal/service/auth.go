package service

import (
	"context"
	"fmt"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	r repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *AuthService {
	return &AuthService{r: r}
}

func (s *AuthService) Register(ctx context.Context, dto model.UserAuthDTO) error {
	now := time.Now().UTC()

	pass, err := helpers.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	user := &model.User{
		ID:        uuid.New(),
		Username:  dto.Username,
		Password:  pass,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.r.Register(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, dto model.UserAuthDTO) (string, error) {
	user, err := s.r.GetUserByUsername(ctx, dto.Username)

	if err != nil || !helpers.CheckPasswordHash(dto.Password, user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
