package middlewares

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Placeholder for authentication logic
		next.ServeHTTP(w, r)
	})
}