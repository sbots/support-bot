package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	botTPlatformTelegram = "telegram"
	botPlatformViber     = "viber"
)

type Bot struct {
	ID        string    `json:"id" db:"id"`
	Token     string    `json:"token" db:"token"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewBot(token, platform string) *Bot {
	return &Bot{ID: uuid.NewV4().String(), Token: token, Type: platform}
}

func (b Bot) IsTelegramBot() bool {
	return b.Type == botTPlatformTelegram
}

func (b Bot) IsViberBot() bool {
	return b.Type == botPlatformViber
}