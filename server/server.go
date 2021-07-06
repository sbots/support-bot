package server

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"support-bot/models"
)

type Server struct {
	http   *http.Server
	tg     telegram
	domain string
	repo   repo
}

type repo interface {
	GetBot(id string) (*models.Bot, error)
	CreateBot(bot *models.Bot) (*models.Bot, error)
}

type telegram interface {
	ConnectNewBot(token, path string) error
	SendMessage(msg *models.Message, token string) error
}

const (
	telegramEndpoint       = "/bots/telegram/"
	telegramNewBotEndpoint = telegramEndpoint + "new/"
	telegramSendMessage    = "send/"
)

func New(addr, domain string, tg telegram, r repo) *Server {
	s := &Server{
		http:   &http.Server{Addr: addr},
		tg:     tg,
		domain: domain,
		repo:   r,
	}
	s.http.Handler = s.router()
	return s
}

func (s Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		err := s.http.Shutdown(ctx)
		if err != nil {
			log.Println("http service shutdown (", err, ")")
		}
		log.Println("service gracefully stopped")
	}()
	log.Print("starting server")
	return s.http.ListenAndServe()
}

func (s *Server) router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(telegramEndpoint+"{bot}", s.telegramHandler)
	router.HandleFunc(telegramNewBotEndpoint, s.newBot)
	router.HandleFunc(telegramSendMessage, s.telegramHandler)

	return router
}

func (s *Server) getEndpointForBot(id string) string {
	return telegramEndpoint + id
}
