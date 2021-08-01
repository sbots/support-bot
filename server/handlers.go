package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"support-bot/models"
	"support-bot/repository/telegram"
)

func (s *Server) webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	var update telegram.Update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(update.Message.Chat, update.Message.Text)
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
	defer r.Body.Close()

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := uuid.NewV4().String()

	bot := models.NewBot(id, data.Token, data.Type)
	if _, err := s.repo.CreateBot(bot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.connectBot(bot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.Marshal(bot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(bytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s Server) send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

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
		http.NotFound(w, r)
		return
	}
	if err := s.sendMessage(bot, msg, req.Receiver); err != nil {
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
