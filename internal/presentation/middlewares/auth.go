package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/infrastructure/auth"
)

type contextKey string

const UserContextKey contextKey = "user"

type UserContext struct {
	UserUUID   uuid.UUID
	Privileges int
}

type AuthMiddleware struct {
	TokenManager auth.TokenManager
}

func (m *AuthMiddleware) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("authentication")
		if err != nil {
			http.Error(w, "Auth token required", http.StatusUnauthorized)
			return
		}

		userUUID, privileges, err := m.TokenManager.ParseAuthToken(cookie.Value)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, UserContext{
			UserUUID:   userUUID,
			Privileges: privileges,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
