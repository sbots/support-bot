package telegram

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/url"
)

type Bot struct {
	client *tg.BotAPI
}

func AddNewBot(token string) (*Bot, error) {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("adding new bot, %w", err)
	}
	return &Bot{bot}, nil
}

func (b Bot) SetNewWebhook(path string) error {
	u, err := url.Parse(path)
	if err != nil{
		return err
	}
	_, err = b.client.SetWebhook(tg.WebhookConfig{
		URL:            u,
		Certificate:    nil,
		MaxConnections: 1,
	})
	if err != nil {
		return fmt.Errorf("setting webhook, %w", err)
	}
	return nil
}
