package middleware

import (
	"context"
	"log"
	"mzda/internal/utils"
	"net/http"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/middleware/auth/JWTAuth"
		token := r.Header.Get("Authorization")
		jwt, err := utils.NewJWT(token)
		if err != nil {
			log.Printf("%s %v", fn, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "jwt", jwt)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
