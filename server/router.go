package server

import "github.com/gorilla/mux"

const (
	telegramEndpoint       = "/bots/telegram/"
	telegramNewBotEndpoint = telegramEndpoint + "new/"
	telegramSendMessage    = telegramEndpoint + "send/"

	viberEndpoint       = "/bots/viber/"
	viberNewBotEndpoint = viberEndpoint + "new"
	viberSendMessage    = viberEndpoint + "send"

	signUpTenantEndpoint = "tenants/new"
	signUpUserEndpoint   = "users/new"

	signInEndpoint   = "/login"
	userInfoEndpoint = "/user"
)

func (s *Server) router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(telegramEndpoint+"{bot}", s.addJWTTokenToContext(s.webhook))
	router.HandleFunc(telegramNewBotEndpoint, s.addJWTTokenToContext(s.newBot))
	router.HandleFunc(telegramSendMessage+"{bot}", s.addJWTTokenToContext(s.send))

	router.HandleFunc(viberEndpoint+"{bot}", s.addJWTTokenToContext(s.webhook))
	router.HandleFunc(viberNewBotEndpoint, s.addJWTTokenToContext(s.newBot))
	router.HandleFunc(viberSendMessage+"{bot}", s.addJWTTokenToContext(s.send))

	router.HandleFunc(signUpTenantEndpoint, s.newTenant)
	router.HandleFunc(signUpUserEndpoint, s.signUp)
	router.HandleFunc(signInEndpoint, s.signIn)

	router.HandleFunc(userInfoEndpoint, s.addJWTTokenToContext(s.getUserInformation))

	return router
}
