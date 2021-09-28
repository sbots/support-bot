package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	http   *http.Server
	domain string
}

func New(addr, domain string, handler http.Handler) *Server {
	s := &Server{
		http:   &http.Server{Addr: addr},
		domain: domain,
	}
	s.http.Handler = handler
	return s
}

func (s Server) Run(ctx context.Context) {
	go func() {
		<-ctx.Done()
		err := s.http.Shutdown(ctx)
		if err != nil {
			log.Println("http service shutdown (", err, ")")
		}
		log.Println("service gracefully stopped")
	}()
	log.Print("starting server")
	if err := s.http.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
