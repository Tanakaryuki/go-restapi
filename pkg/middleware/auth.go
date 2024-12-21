package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Tanakaryuki/go-restapi/pkg/auth"
)

const (
	UsernameKey = "username"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		actualToken := strings.TrimPrefix(token, "Bearer ")
		claims, err := auth.ValidateToken(actualToken)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UsernameKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
