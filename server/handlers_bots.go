package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"support-bot/models"
)

func (s *Server) webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	var objmap map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&objmap)
	defer closeBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println(objmap)
}

func (s *Server) newBot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Token string `json:"token"`
		Type  string `json:"type"`
	}

	decoder := json.NewDecoder(r.Body)
	defer closeBody(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bot := models.NewBot(data.Token, data.Type)
	if err := s.connectBot(bot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := s.repo.UpsertBot(bot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := prepareResponse(w, bot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s Server) send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer closeBody(r.Body)

	var req struct {
		Receiver string `json:"receiver"`
		Text     string `json:"text"`
	}

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msg := models.NewMessage("", req.Text)
	bot, err := s.repo.GetBot(mux.Vars(r)["bot"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = s.sendMessage(bot, msg, req.Receiver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// TODO: move this check to repository
func (s Server) connectBot(bot *models.Bot) error {
	if bot.IsTelegramBot() {
		path := s.domain + s.getEndpointForTgBot(bot.ID)
		return s.tg.ConnectNewBot(bot.Token, path)
	}
	if bot.IsViberBot() {
		path := s.domain + s.getEndpointForVbBot(bot.ID)
		return s.vb.ConnectNewBot(bot.Token, path)
	}
	return fmt.Errorf("unsupported platform")
}

func (s Server) sendMessage(bot *models.Bot, msg *models.Message, receiver string) error {
	if bot.IsTelegramBot() {
		return s.tg.SendMessage(msg, bot.Token, receiver)
	}
	if bot.IsViberBot() {
		return s.vb.SendMessage(msg, bot.Token, receiver)
	}
	return fmt.Errorf("unsupported platform")
}
