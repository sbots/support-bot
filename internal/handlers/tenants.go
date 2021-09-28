package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"support-bot/internal/models"
)

func (c *controller) newTenant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	var form models.NewCompanyForm
	if ok := decodeRequest(r.Body, &form); !ok {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	tenant, err := c.service.NewTenant(form.Name)
	if err != nil {
		log.WithError(err).Debug("creating new tenant")
		http.Error(w, "", errorToHttpCode(err))
		return
	}
	if ok := prepareResponse(w, tenant); !ok {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
