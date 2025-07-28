package service

import (
	"context"
	"errors"
	"fmt"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	r  repository.AuthRepository
	es *EmailService
}

func NewAuthService(r repository.AuthRepository, es *EmailService) *AuthService {
	return &AuthService{r: r, es: es}
}

func (s *AuthService) Register(ctx context.Context, dto model.UserRegisterDTO) (*model.UserRegisterResponseDTO, error) {
	now := time.Now().UTC()

	pass, err := helpers.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        uuid.New(),
		Username:  dto.Username,
		Password:  pass,
		Email:     dto.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.r.Register(ctx, user); err != nil {
		return nil, err
	}

	code := generateCode()
	if err := s.r.SetCode(ctx, user.ID, code); err != nil {
		return nil, err
	}

	if err = s.es.SendEmail(dto.Email,
		"ContentLog Verification",
		"Your verification code for ContentLog!",
		fmt.Sprintf("Your code <b>%s</b>", code)); err != nil {
		return nil, err
	}

	return &model.UserRegisterResponseDTO{ID: user.ID}, nil
}

func (s *AuthService) Login(ctx context.Context, dto model.UserLoginDTO) (string, error) {
	user, err := s.r.GetUserByEmail(ctx, dto.Email)

	if err != nil || !helpers.CheckPasswordHash(dto.Password, user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	if !user.Confirmed {
		return "", fmt.Errorf("email is not confirmed")
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Verify(ctx context.Context, id uuid.UUID, dto *model.UserVerifyDTO) error {
	res, err := s.r.GetCode(ctx, id)
	if err != nil {
		return err
	}

	if res != dto.Code {
		return errors.New("wrong code")
	}

	if err := s.r.Verify(ctx, id, time.Now().UTC()); err != nil {
		return err
	}

	return nil
}

func generateCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}

	return code
}
