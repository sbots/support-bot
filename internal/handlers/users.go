package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c controller) getUserInformation(w http.ResponseWriter, r *http.Request) {
	user, err := c.service.GetUserInformation(r.Context())
	if err != nil {
		log.WithError(err).Debug("getting user info")
		http.Error(w, "getting user information", errorToHttpCode(err))
	}

	if ok := prepareResponse(w, user); !ok {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
