package middleware

import (
	"context"
	"net/http"

	"energy_tk/pkg/render"
)

func AppAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userToken := r.Header.Get("Authorization")
			if userToken == "" {
				render.Json(w, http.StatusBadRequest, "token is required")
				return
			}
			ctx := context.WithValue(r.Context(), "token", userToken)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserToken(ctx context.Context) string {
	return ctx.Value("token").(string)
}
