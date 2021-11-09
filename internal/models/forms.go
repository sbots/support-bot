package models

type SignUpForm struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type SignInForm struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type NewCompanyForm struct {
	Name string `json:"name"`
}

type ChatMessage struct {
	Token string `json:"token"` // user id
	To    string `json:"to"`    // bot id
	Text  string `json:"text"`
}

type BotMessage struct {
	Bot  string `json:"bot"`  // bot id
	Chat string `json:"from"` // chat id
	To   string `json:"to"`   // user id
	Text string `json:"text"`
}
