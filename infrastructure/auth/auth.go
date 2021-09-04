package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"support-bot/models"
	"time"
)

type contextKey int

const serviceTokenContextKey contextKey = 1

type Authenticator struct {
	signingKey []byte
	expiration time.Duration
	issuer     string
}

func NewAuthenticator(key, issuer string, expiration time.Duration) (*Authenticator, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("signing key is required")
	}
	return &Authenticator{
		signingKey: []byte(key),
		expiration: expiration,
		issuer:     issuer,
	}, nil
}

func (a *Authenticator) GetToken(user *models.User) (string, error) {
	unsignedToken := newUnsignedToken(user.ID, user.Company, a.issuer, a.expiration)
	return a.signToken(unsignedToken)
}

func (a *Authenticator) ParseToken(tokenString string) (JWTToken, error) {
	token, err := a.parseToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}
	err = token.Valid()
	if err != nil {
		return nil, err
	}
	return token, err
}

func (a *Authenticator) SetServiceTokenToContext(ctx context.Context, serviceToken JWTToken) context.Context {
	return context.WithValue(ctx, serviceTokenContextKey, serviceToken)
}

func (a *Authenticator) GetServiceTokenFromContext(ctx context.Context) JWTToken {
	return ctx.Value(serviceTokenContextKey).(*Token)
}

func (a *Authenticator) signToken(t *Token) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, t).SignedString(a.signingKey)
}

func (a *Authenticator) parseToken(tokenString string) (*Token, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return a.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	token, ok := jwtToken.Claims.(*Token)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return token, err
}
