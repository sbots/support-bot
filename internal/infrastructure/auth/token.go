package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTToken interface {
	GetUserID() string
	GetTenantID() string
}

type Token struct {
	jwt.StandardClaims
	CompanyID string `json:"company_id"`
	UserID    string `json:"user_id"`
}

func newUnsignedToken(userID, companyID, issuer string, expiresIn time.Duration) *Token {
	return &Token{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().UTC().Add(expiresIn).Unix(),
		},
		CompanyID: companyID,
		UserID:    userID,
	}
}

func (t *Token) GetUserID() string {
	return t.UserID
}

func (t *Token) GetTenantID() string {
	return t.CompanyID
}
