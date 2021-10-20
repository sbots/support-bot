package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"support-bot/internal/models"
)

func (c controller) signUp(w http.ResponseWriter, r *http.Request) {
	var form models.SignUpForm
	if ok := decodeRequest(r.Body, &form); !ok {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	token, err := c.service.UserSignUp(&form)
	if err != nil {
		log.WithError(err).Debug("signing up user")
		http.Error(w, "", errorToHttpCode(err))
		return
	}
	if ok := prepareResponse(w, &token); !ok {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *controller) signIn(w http.ResponseWriter, r *http.Request) {
	var form models.SignInForm
	if ok := decodeRequest(r.Body, &form); !ok {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	token, err := c.service.UserSignIn(&form)
	if err != nil {
		log.WithError(err).Debug("signing in user")
		http.Error(w, "", errorToHttpCode(err))
		return
	}

	if ok := prepareResponse(w, &token); !ok {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
