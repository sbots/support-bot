package models

type Bot struct {
	ID string `json:"id" gorm:"id"`
	Token string `json:"token" gorm:"token"`
}

func NewBot(id, token string) *Bot{
	return &Bot{
		ID:   id,
		Token: token,
	}
}