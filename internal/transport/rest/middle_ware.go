package rest

import (
	"github.com/Feinot/simple_bank/internal/modules/token"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

const (
	authHeaderKey    = "Authorization"
	authHeaderPrefix = "Bearer"
)

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(authHeaderKey)
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if headerParts[0] != authHeaderPrefix {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		signingKey := []byte(viper.GetString("auth.signing_key"))
		if _, err := token.ParseToken(headerParts[1], signingKey); err != nil {
			status := http.StatusBadRequest
			if token.IsErrInvalidAccessToken(err) {
				status = http.StatusUnauthorized
			}
			w.WriteHeader(status)
			return
		}
		next.ServeHTTP(w, r)
	})
}
