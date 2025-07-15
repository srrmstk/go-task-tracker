package logger

import (
	"log/slog"
	"net/http"
)

func LoggerMiddleware(log *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request", "method", r.Method, "url", r.URL.String())
		next(w, r)
	})
}
