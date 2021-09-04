package server

import "net/http"

func (s *Server) getUserInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	token := s.authenticator.GetServiceTokenFromContext(r.Context())
	if token == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := s.repo.GetUserByID(token.GetUserID())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = prepareResponse(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
