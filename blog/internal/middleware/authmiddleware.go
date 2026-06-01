package middleware

import (
	"blog/internal/auth"
	"blog/internal/model"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type contextKey string

const claimsContextKey contextKey = "authClaims"

func Authenticate(secret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			writeAuthError(w, http.StatusUnauthorized, "missing authorization header")
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			writeAuthError(w, http.StatusUnauthorized, "invalid authorization header")
			return
		}

		claims, err := auth.ParseToken(parts[1], secret)
		if err != nil {
			writeAuthError(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), claimsContextKey, claims)
		next(w, r.WithContext(ctx))
	}
}

func RequireRole(role model.Role, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := ClaimsFromContext(r.Context())
		if !ok {
			writeAuthError(w, http.StatusUnauthorized, "missing authentication context")
			return
		}

		if claims.Role != role {
			writeAuthError(w, http.StatusForbidden, "insufficient permissions")
			return
		}

		next(w, r)
	}
}

func ClaimsFromContext(ctx context.Context) (auth.Claims, bool) {
	claims, ok := ctx.Value(claimsContextKey).(auth.Claims)
	return claims, ok
}

func writeAuthError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
