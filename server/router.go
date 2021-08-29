package server

import "github.com/gorilla/mux"

const (
	telegramEndpoint       = "/bots/telegram/"
	telegramNewBotEndpoint = telegramEndpoint + "new/"
	telegramSendMessage    = telegramEndpoint + "send/"

	viberEndpoint       = "/bots/viber/"
	viberNewBotEndpoint = viberEndpoint + "new/"
	viberSendMessage    = viberEndpoint + "/send"

	signInTenantEndpoint = "/tenants/new"
	signInUserEndpoint   = "/users/new"

	loginEndpoint = "/login"
)

func (s *Server) router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(telegramEndpoint+"{bot}", s.webhook)
	router.HandleFunc(telegramNewBotEndpoint, s.newBot)
	router.HandleFunc(telegramSendMessage+"{bot}", s.send)

	router.HandleFunc(viberEndpoint+"{bot}", s.webhook)
	router.HandleFunc(viberNewBotEndpoint, s.newBot)
	router.HandleFunc(viberSendMessage+"{bot}", s.send)

	router.HandleFunc(signInTenantEndpoint, s.newTenant)
	router.HandleFunc(signInUserEndpoint, s.newUser)

	return router
}
