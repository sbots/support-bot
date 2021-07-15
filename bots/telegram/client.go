package telegram

import (
	"errors"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/url"
	"support-bot/models"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) SendMessage(msg *models.Message, token string) error {
	if msg == nil {
		return errors.New("message input is missing")
	}
	m := tg.NewMessage(msg.Chat.ID, msg.Text)
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return errors.New("get bot by token")
	}
	resp, err := bot.Send(m)
	fmt.Println(resp)
	if err != nil {
		return fmt.Errorf("sending message %w", err)
	}
	return nil
}

func (c Client) ConnectNewBot(token, path string) error {
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
