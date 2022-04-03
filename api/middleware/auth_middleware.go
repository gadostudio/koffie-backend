package middleware

import (
	"context"
	"github.com/shaderboi/koffie-backend/api/settings"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		firebaseAuth := settings.SetupFirebase(context.Background())

		rawToken := r.Header.Get("Authorization")

		token := strings.Split(rawToken, " ")[1]

		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err := firebaseAuth.VerifyIDToken(context.Background(), token)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

		next.ServeHTTP(w, r)
	})
}
