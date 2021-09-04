package auth_test

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"support-bot/infrastructure/auth"
	"support-bot/models"
	"testing"
	"time"
)

const (
	testSigningKey = "secret"
	testIssuer     = "support-bot-platform-test"
)

func TestAuthenticator_GetToken(t *testing.T) {
	user := &models.User{
		ID:      "13",
		Company: "12",
	}
	a, err := auth.NewAuthenticator(testSigningKey, testIssuer, time.Minute*5)
	assert.NoError(t, err)
	_, err = a.GetToken(user)
	assert.NoError(t, err)
}

func TestAuthenticator_ParseToken(t *testing.T) {
	test := struct {
		name   string
		key    string
		issuer string
		user   *models.User
		token  *auth.Token
	}{
		name:   "success",
		key:    testSigningKey,
		issuer: testIssuer,
		user: &models.User{
			ID:      "13",
			Company: "12",
		},
		token: &auth.Token{
			StandardClaims: jwt.StandardClaims{
				Issuer: testIssuer,
			},
			CompanyID: "12",
			UserID:    "13",
		},
	}
	a, err := auth.NewAuthenticator(test.key, testIssuer, time.Minute*5)
	assert.NoError(t, err)

	tok, err := a.GetToken(test.user)
	assert.NoError(t, err)

	got, err := a.ParseToken(tok)
	assert.NoError(t, err)
	assert.Equal(t, test.token.UserID, got.GetUserID())
	assert.Equal(t, test.token.CompanyID, got.GetTenantID())
}
