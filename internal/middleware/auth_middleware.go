package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	jwtutil "github.com/sachin-gautam/go-crud-api/internal/utils/jwt"
)

// AuthMiddleware checks the JWT in the Authorization header
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtutil.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		slog.Info("User authenticated", slog.String("username", claims.Username))

		next.ServeHTTP(w, r)
	})
}
