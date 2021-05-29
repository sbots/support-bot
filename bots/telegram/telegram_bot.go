package telegram

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/url"
)

type Bot struct {
	client *tg.BotAPI
}

func AddNewBot(token string) (*Bot, error) {
	fmt.Println(token)
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("adding new bot, %w", err)
	}
	return &Bot{bot}, nil
}

func (b Bot) SetNewWebhook(path string) error {
	rsp, err := b.client.SetWebhook(tg.WebhookConfig{
		URL:            &url.URL{Path: path},
		Certificate:    nil,
		MaxConnections: 1,
	})
	if err != nil {
		return fmt.Errorf("setting webhook, %w", err)
	}
	log.Println(rsp)
	return nil
}
