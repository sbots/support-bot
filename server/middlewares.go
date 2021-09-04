package server

import (
	"context"
	"net/http"
)

const authHeaderKey = "authorization"

func (s *Server) addJWTTokenToContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationValue := r.Header.Get(authHeaderKey)
		if authorizationValue != "" {
			token, err := s.authenticator.ParseToken(authorizationValue)
			if err == nil {
				ctx := s.authenticator.SetServiceTokenToContext(context.Background(), token)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	}
}
