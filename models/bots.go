package models

const (
	botTPlatformTelegram = "telegram"
	botPlatformViber     = "viber"
)

type Bot struct {
	ID    string `json:"id" gorm:"id"`
	Token string `json:"token" gorm:"token"`
	Type  string `json:"type" gorm:"type"`
}

func NewBot(id, token, platform string) *Bot {
	return &Bot{ID: id, Token: token, Type: platform}
}

func (b Bot) IsTelegramBot() bool {
	return b.Type == botTPlatformTelegram
}

func (b Bot) IsViberBot() bool {
	return b.Type == botPlatformViber
}
