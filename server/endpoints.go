package server

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
