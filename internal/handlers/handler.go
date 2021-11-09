package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"support-bot/internal/infrastructure/auth"
	"support-bot/internal/infrastructure/errors"
	"support-bot/internal/models"
)

type controller struct {
	service        service
	auth           auth.Authenticator
	productionMode bool
	upgrader       websocket.Upgrader
}

type service interface {
	UserSignUp(form *models.SignUpForm) (string, error)
	UserSignIn(form *models.SignInForm) (string, error)
	GetUserInformation(ctx context.Context) (*models.User, error)

	NewTenant(name string) (*models.Tenant, error)
}

func NewHandler(s service, auth auth.Authenticator, testMode bool) http.Handler {
	c := &controller{service: s, auth: auth, productionMode: testMode}
	c.setupWSUpgrader()
	return c.buildHandler()
}

func prepareResponse(w http.ResponseWriter, rsp interface{}) bool {
	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.WithError(err).Println("preparing response")
		return false
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.WithError(err).Println("preparing response")
		return false
	}
	w.Header().Set("Content-Type", "application/json")
	return true
}

func closeBody(b io.ReadCloser) {
	if err := b.Close(); err != nil {
		log.WithError(err).Println("closing response body")
	}
}

func decodeRequest(body io.ReadCloser, res interface{}) bool {
	decoder := json.NewDecoder(body)
	defer closeBody(body)

	if err := decoder.Decode(res); err != nil {
		log.Errorf("decoding request: %s", err.Error())
		return false
	}
	return true
}

func errorToHttpCode(err error) int {
	if errors.Is(err, errors.NotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, errors.AccessDenied) {
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}
