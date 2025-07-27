package service

import (
	"context"
	"fmt"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gopkg.in/mail.v2"
)

type AuthService struct {
	r repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *AuthService {
	return &AuthService{r: r}
}

func (s *AuthService) Register(ctx context.Context, dto model.UserRegisterDTO) error {
	now := time.Now().UTC()

	pass, err := helpers.HashPassword(dto.Password)
	if err != nil {
		return err
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
		return err
	}

	err = SendEmail(dto.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, dto model.UserLoginDTO) (string, error) {
	user, err := s.r.GetUserByEmail(ctx, dto.Email)

	if err != nil || !helpers.CheckPasswordHash(dto.Password, user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func SendEmail(to string) error {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetAddressHeader("From", from, "ContentLog Verification")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your verification code for ContentLog!")
	m.SetBody("text/html", fmt.Sprintf("Your code <b>%d</b>", 123456))
	d := mail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
