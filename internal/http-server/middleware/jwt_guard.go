package middleware

import (
	"context"
	"go-task-tracker/internal/helpers"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func JwtGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			helpers.JsonError(w, "wrong Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			slog.Debug(err.Error())
			helpers.JsonError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok {
			if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
				helpers.JsonError(w, "token sexpired", http.StatusUnauthorized)
				return

			}
			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}

		helpers.JsonError(w, "invalid token", http.StatusUnauthorized)
	})
}
