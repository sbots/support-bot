package models

import uuid "github.com/satori/go.uuid"

const (
	botTPlatformTelegram = "telegram"
	botPlatformViber     = "viber"
)

type Bot struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	Type  string `json:"type"`
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
