package telegram

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/url"
	"strconv"
	"support-bot/models"
)

const (
	// APIEndpoint is the endpoint for all API methods,
	// with formatting for Sprintf.
	APIEndpoint = "https://api.telegram.org/bot%s/%s"
	// FileEndpoint is the endpoint for downloading a file from Telegram.
	FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) SendMessage(msg *models.Message, token, receiver string) error {
	if msg == nil {
		return fmt.Errorf("message input is missing")
	}

	chat, err := strconv.ParseInt(receiver, 0, 32)
	if err != nil {
		return fmt.Errorf("parsing chat id: %w", err)
	}

	m := tg.NewMessage(chat, msg.Text)
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return fmt.Errorf("get bot by token")
	}
	resp, err := bot.Send(m)
	log.Println(resp)
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
