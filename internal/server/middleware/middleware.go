package middleware

import (
	"net/http"

	"golang.org/x/exp/slog"
)

func Logger(log *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request recieved",
			slog.String("request uri", r.RequestURI),
			slog.String("user-agent", r.UserAgent()),
			slog.String("method", r.Method))

		next(w, r)
	}
}
