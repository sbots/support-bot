package server

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"support-bot/errors"
	"support-bot/models"
)

func (s *Server) newTenant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	defer closeBody(r.Body)

	if err := decoder.Decode(&data); err != nil {
		log.Errorf("decoding request: %s", err.Error())
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	tenant := models.NewTenant(data.Name)
	if err := s.repo.UpsertTenant(tenant); err != nil {
		log.Errorf("upserting tenant: %s", err.Error())
		http.Error(w, "creating tenant failed", http.StatusUnprocessableEntity)
		return
	}

	if err := prepareResponse(w, tenant); err != nil {
		log.Errorf("preparing response: %s", err.Error())
		http.Error(w, errors.InternalServerError, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Password string `json:"password"`
		Company  string `json:"company"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}

	decoder := json.NewDecoder(r.Body)
	defer closeBody(r.Body)

	if err := decoder.Decode(&data); err != nil {
		log.Errorf("decoding request: %s", err.Error())
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if yes, err := s.repo.TenantHasSuperuser(data.Company); err != nil {
		log.Errorf("checking for superuser failed: %s", err.Error())
		http.Error(w, errors.InternalServerError, http.StatusInternalServerError)
		return
	} else if yes {
		http.Error(w, "currently only one user per company is allowed", http.StatusConflict)
		return
	}

	user, err := models.NewUser(data.Name, data.Surname, data.Password, data.Company, data.Email, data.Phone)
	if err != nil {
		log.Errorf("creating user failed: %s", err.Error())
		http.Error(w, errors.InternalServerError, http.StatusUnprocessableEntity)
		return
	}

	err = s.repo.UpsertUser(user)
	if err != nil {
		log.Errorf("upserting user failed: %s", err.Error())
		http.Error(w, errors.InternalServerError, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	defer closeBody(r.Body)

	if err := decoder.Decode(&data); err != nil {
		log.Println(fmt.Errorf("decoding request %w", err))
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	user, err := s.repo.GetUserByEmail(data.Email)
	if err != nil {
		if err.Error() == errors.NotFound {
			http.Error(w, "not registered", http.StatusNotFound)
			return
		}
		log.Println(fmt.Errorf("getting user by email %w", err))
		http.Error(w, errors.InternalServerError, http.StatusInternalServerError)
		return
	}

	if !user.ValidPassword(data.Password) {
		http.Error(w, "wrong password", http.StatusForbidden)
		return
	}

	token, err := s.authenticator.GetToken(user)
	if err != nil {
		log.Errorf("getting authorisation token: %s", err.Error())
		http.Error(w, errors.InternalServerError, http.StatusInternalServerError)
		return
	}

	if err := prepareResponse(w, token); err != nil {
		log.Errorf("preparing response: %s", err.Error())
		http.Error(w, errors.InternalServerError, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
