package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	models2 "support-bot/api/models"
	"time"
)

type Authenticator struct {
	signingKey []byte
}

func NewAuthenticator(key []byte) (*Authenticator, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("signing key is required")
	}
	return &Authenticator{
		signingKey: key,
	}, nil
}

func (a *Authenticator) GetJWT(user *models2.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // nolint

	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["iss"] = "support-bot-platform-test"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", fmt.Errorf("signing token string: %w", err)
	}
	return tokenString, nil
}
