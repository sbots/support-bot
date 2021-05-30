package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"support-bot/models"
)

type Server struct {
	http             *http.Server
	telegramEndpoint string
}

const telegramEndpoint = "/bots/telegram/"

func New(addr string) *Server {
	s := &Server{
		http:             &http.Server{Addr: addr},
		telegramEndpoint: telegramEndpoint,
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
		log.Println(errors.New("wrong HTTP method required POST"))
	}

	var update models.Update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Println(err)
	}

	log.Println(update.Message.Text)
}

func (s *Server) GetEndpointForBot(id string) string {
	return s.telegramEndpoint + id
}

func getBot(r *http.Request) string {
	return mux.Vars(r)["bot"]
}
