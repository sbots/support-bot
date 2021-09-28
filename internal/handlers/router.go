package handlers

import (
	"github.com/gorilla/mux"
)

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

func (c *controller) buildHandler() *mux.Router {
	router := mux.NewRouter()

	//router.HandleFunc(telegramEndpoint+"{bot}", c.AddJWTTokenToContext(c.webhook))
	//router.HandleFunc(telegramSendMessage+"{bot}", c.AddJWTTokenToContext(c.send))
	//
	//router.HandleFunc(viberEndpoint+"{bot}", c.AddJWTTokenToContext(c.webhook))
	//router.HandleFunc(viberSendMessage+"{bot}", c.AddJWTTokenToContext(c.send))

	router.HandleFunc(signUpTenantEndpoint, c.newTenant)
	router.HandleFunc(signUpUserEndpoint, c.signUp)
	router.HandleFunc(signInEndpoint, c.signIn)

	//router.HandleFunc(newBotEndpoint, c.newBot)

	router.HandleFunc(userInfoEndpoint, c.AddJWTTokenToContext(c.getUserInformation))
	//router.HandleFunc(chat, c.Chat)

	return router
}
