package handlers

import (
	"context"
	"net/http"
)

const authHeaderKey = "authorization"

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

func (c *controller) allowCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !c.productionMode {
			r.Header.Set("Access-Control-Allow-Origin", "*")
		}
		next.ServeHTTP(w, r)
	}
}
