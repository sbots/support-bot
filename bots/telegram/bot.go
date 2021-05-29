package telegram

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/url"
)

func ConnectNewBot(token, path string) error {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return fmt.Errorf("adding new bot, %w", err)
	}
	u, err := url.Parse(path)
	if err != nil {
		return err
	}
	_, err = bot.SetWebhook(tg.WebhookConfig{
		URL:            u,
		Certificate:    nil,
		MaxConnections: 1,
	})
	if err != nil {
		return fmt.Errorf("setting webhook, %w", err)
	}
	return nil
}
