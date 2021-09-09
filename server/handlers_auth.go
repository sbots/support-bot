package server

import (
	"encoding/json"
	"net/http"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tenant := models.NewTenant(data.Name)
	if err := s.repo.UpsertTenant(tenant); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := prepareResponse(w, tenant); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if yes, err := s.repo.TenantHasSuperuser(data.Company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if yes {
		http.Error(w, "only one user currently allowed", http.StatusConflict)
		return
	}

	user, err := models.NewUser(data.Name, data.Surname, data.Password, data.Company, data.Email, data.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = s.repo.UpsertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := s.repo.GetUserByEmail(data.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !user.ValidPassword(data.Password) {
		http.Error(w, "wrong password", http.StatusForbidden)
	}

	token, err := s.authenticator.GetToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := prepareResponse(w, token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
