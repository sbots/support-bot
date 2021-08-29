package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"support-bot/models"
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

func (a *Authenticator) GetJWT(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // nolint

	claims["authorized"] = true
	claims["name"] = user.Name
	claims["surname"] = user.Surname
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", fmt.Errorf("signing token string: %w", err)
	}
	return tokenString, nil
}
