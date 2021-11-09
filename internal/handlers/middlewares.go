package handlers

import (
	"context"
	"net/http"
)

const (
	authHeaderKey = "authorization"
	tokenField    = "token"
)

func (c *controller) addJWTTokenToContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationValue := r.Header.Get(authHeaderKey)
		if authorizationValue != "" {
			token, err := c.auth.ParseToken(authorizationValue)
			if err == nil {
				ctx := c.auth.SetServiceTokenToContext(context.Background(), token)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	}
}

func (c *controller) validateRequest(r *http.Request) bool {
	token, err := c.auth.ParseToken(r.Header.Get(tokenField))
	if err != nil {
		return false
	}
	err = token.Valid()
	return err == nil
}

func (c *controller) allowCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !c.productionMode {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}
		next.ServeHTTP(w, r)
	}
}
