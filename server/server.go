package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	router.HandleFunc(tgPath, s.handler)
	return router
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	bot := getBot(r)
	if _, err := fmt.Fprintf(w, bot); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) GetEndpointForBot(id string) string {
	return s.telegramEndpoint + id
}

func getBot(r *http.Request) string {
	return mux.Vars(r)["bot"]
}
