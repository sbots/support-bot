package server

import "github.com/gorilla/mux"

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
