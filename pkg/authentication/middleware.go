package authentication

import (
	"context"
	"net/http"
	"strconv"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing token",
					http.StatusUnauthorized)
				return
			}

			idStr, err := ParseToken(secret, authHeader)
			if err != nil {
				http.Error(w, "Invalid token",
					http.StatusUnauthorized)
				return
			}

			userID, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid token",
					http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, uint(userID))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) uint {
	id, _ := ctx.Value(userIDKey).(uint)
	return id
}
