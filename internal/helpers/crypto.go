package helpers

import (
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(id uuid.UUID) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	slog.Debug(secret)
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(secret)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": id,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		})
	s, err := t.SignedString(key)

	return s, err
}
