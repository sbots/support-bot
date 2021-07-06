package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"support-bot/models"
)

func (s *Server) telegramHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusInternalServerError)
		return
	}

	var update models.Update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(update.Message.Text)
}

func (s *Server) newBot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusInternalServerError)
		return
	}

	bot := models.NewBot(uuid.NewV4().String(), mux.Vars(r)["token"])
	if _, err := s.repo.CreateBot(bot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := s.domain + s.getEndpointForBot(bot.ID)
	if err := s.tg.ConnectNewBot(bot.Token, path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(bot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s Server) send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var msg models.Message
	err := decoder.Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.tg.SendMessage(&msg, mux.Vars(r)["token"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
