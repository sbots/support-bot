package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
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

	//router.HandleFunc(telegramEndpoint+"{bot}", c.addJWTTokenToContext(c.webhook))
	//router.HandleFunc(telegramSendMessage+"{bot}", c.addJWTTokenToContext(c.send))
	//
	//router.HandleFunc(viberEndpoint+"{bot}", c.addJWTTokenToContext(c.webhook))
	//router.HandleFunc(viberSendMessage+"{bot}", c.addJWTTokenToContext(c.send))

	router.HandleFunc(signUpTenantEndpoint, c.allowCORS(c.newTenant)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc(signUpUserEndpoint, c.allowCORS(c.signUp)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc(signInEndpoint, c.allowCORS(c.signIn)).Methods(http.MethodPost, http.MethodOptions)

	//router.HandleFunc(newBotEndpoint, c.newBot)

	router.HandleFunc(userInfoEndpoint, c.allowCORS(c.addJWTTokenToContext(c.getUserInformation))).
		Methods(http.MethodPost, http.MethodOptions)
	//router.HandleFunc(chat, c.Chat)
	return router
}
