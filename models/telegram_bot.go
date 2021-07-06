package models

import uuid "github.com/satori/go.uuid"

type Bot struct {
	ID    string `json:"id" gorm:"id"`
	Token string `json:"token" gorm:"token"`
}

func (b *Bot) SetUUID() {
	b.ID = uuid.NewV4().String()
}

func NewBot(id, token string) *Bot {
	return &Bot{
		ID:    id,
		Token: token,
	}
}
