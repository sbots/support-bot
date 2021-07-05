package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"support-bot/models"
)

type Server struct {
	http             *http.Server
	telegramEndpoint string
	tg               telegram
	domain           string
	repo             repo
}

type repo interface {
	GetBot(id string) (*models.Bot,error)
	CreateBot(bot *models.Bot) (*models.Bot, error)
}

type telegram interface {
	ConnectNewBot(token, path string) error
	SendMessage(msg *models.Message, token string) error
}

const telegramEndpoint = "/bots/telegram/"

func New(addr, domain string, tg telegram, r repo) *Server {
	s := &Server{
		http:             &http.Server{Addr: addr},
		telegramEndpoint: telegramEndpoint,
		tg:               tg,
		domain:           domain,
		repo:             r,
	}
	s.http.Handler = s.router()
	return s
}

func (s Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		err := s.http.Shutdown(ctx)
		if err != nil {
			log.Print("http service shutdown (", err, ")")
		}
	}()
	log.Print("starting server")
	return s.http.ListenAndServe()
}

func (s *Server) router() *mux.Router {
	router := mux.NewRouter()
	tgPath := s.telegramEndpoint + "{bot}"
	router.HandleFunc(tgPath, s.telegramHandler)
	return router
}

func (s *Server) telegramHandler(_ http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(fmt.Errorf("wrong HTTP method required POST"))
		return
	}

	var update models.Update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Println(err)
	}

	log.Println(update.Message.Text)
}

func (s *Server) newBot(_ http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(fmt.Errorf("wrong HTTP method required POST"))
		return
	}

	bot := models.NewBot(uuid.NewV4().String(), extractToken(r))
	if _, err := s.repo.CreateBot(bot); err != nil {
		log.Println(fmt.Errorf("error creating bot: %w", err))
		return
	}
	path := s.domain + s.getEndpointForBot(bot.ID)
	if err := s.tg.ConnectNewBot(bot.Token, path); err != nil {
		log.Fatal(err)
	}
	log.Printf("\n new bot{%s} successfully registered with id %s", bot.Token, bot.ID)
}

func (s Server) send(_ http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(fmt.Errorf("wrong HTTP method required POST"))
		return
	}

}

func (s *Server) getEndpointForBot(id string) string {
	return s.telegramEndpoint + id
}

func extractToken(r *http.Request) string {
	return mux.Vars(r)["token"]
}
