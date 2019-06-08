package authmiddleware

import (
	"io"
	"net/http"

	"../../api/authapi"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...check user authoraize
		tokenString := r.Header.Get("token")
		_, err := authapi.ValidateToken(tokenString)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"unauthoraize user"}`)
			return
		}

		next.ServeHTTP(w, r)
	})
}
