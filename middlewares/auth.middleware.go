package middlewares

import (
	"encoding/json"
	"net/http"
)

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			response := map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		next.ServeHTTP(w, r)
	}
}
