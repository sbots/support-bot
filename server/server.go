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
	tg     botsPlatform
	vb     botsPlatform
	domain string
	repo   repo
}

type repo interface {
	GetBot(id string) (*models.Bot, error)
	CreateBot(bot *models.Bot) (*models.Bot, error)
}

type botsPlatform interface {
	ConnectNewBot(token, path string) error
	SendMessage(msg *models.Message, token, receiver string) error
}

const (
	telegramEndpoint       = "/bots/telegram/"
	telegramNewBotEndpoint = telegramEndpoint + "new/"
	telegramSendMessage    = telegramEndpoint + "send/"
	viberEndpoint          = "/bots/viber/"
	viberNewBotEndpoint    = viberEndpoint + "new/"
	viberSendMessage       = viberEndpoint + "/send"
)

func New(addr, domain string, tg, vb botsPlatform, r repo) *Server {
	s := &Server{
		http:   &http.Server{Addr: addr},
		tg:     tg,
		vb:     vb,
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

	router.HandleFunc(telegramEndpoint+"{bot}", s.webhook)
	router.HandleFunc(telegramNewBotEndpoint, s.newBot)
	router.HandleFunc(telegramSendMessage+"{bot}", s.send)

	router.HandleFunc(viberEndpoint+"{bot}", s.webhook)
	router.HandleFunc(viberNewBotEndpoint, s.newBot)
	router.HandleFunc(viberSendMessage+"{bot}", s.send)

	return router
}

func (s *Server) getEndpointForTgBot(id string) string {
	return telegramEndpoint + id
}

func (s *Server) getEndpointForVbBot(id string) string {
	return viberEndpoint + id
}
