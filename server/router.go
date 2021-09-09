package server

import "github.com/gorilla/mux"

const (
	newBotEndpoint = "/bots/new/"

	telegramEndpoint    = "/bots/telegram/"
	telegramSendMessage = telegramEndpoint + "send/"

	viberEndpoint    = "/bots/viber/"
	viberSendMessage = viberEndpoint + "send"

	signUpTenantEndpoint = "/tenants/new"
	signUpUserEndpoint   = "/users/new"

	signInEndpoint   = "/login"
	userInfoEndpoint = "/user"

	chat = "/chat"
)

func (s *Server) router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(telegramEndpoint+"{bot}", s.addJWTTokenToContext(s.webhook))
	router.HandleFunc(telegramSendMessage+"{bot}", s.addJWTTokenToContext(s.send))

	router.HandleFunc(viberEndpoint+"{bot}", s.addJWTTokenToContext(s.webhook))
	router.HandleFunc(viberSendMessage+"{bot}", s.addJWTTokenToContext(s.send))

	router.HandleFunc(signUpTenantEndpoint, s.newTenant)
	router.HandleFunc(signUpUserEndpoint, s.signUp)
	router.HandleFunc(signInEndpoint, s.signIn)

	router.HandleFunc(newBotEndpoint, s.newBot)

	router.HandleFunc(userInfoEndpoint, s.addJWTTokenToContext(s.getUserInformation))
	router.HandleFunc(chat, s.chat)

	return router
}
