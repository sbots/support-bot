package viber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"support-bot/models"
)

const minAPIVersion = 1

const (
	baseURL        = "https://chatapi.viber.com/pa/"
	setWebhookURL  = baseURL + "set_webhook"
	sendMessageURL = baseURL + "send_message"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{
			Transport: LoggingRoundTripper{http.DefaultTransport},
		},
	}
}

func (c Client) ConnectNewBot(token, path string) error {
	sub, b := &SubscriptionRequest{
		Url:        path,
		EventTypes: AllEventsTypes,
		SendName:   true,
		SendPhoto:  true,
	}, new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(sub); err != nil {
		return fmt.Errorf("encoding subsc")
	}
	log.Printf("making request with data %v", sub)
	return c.makeRequest(setWebhookURL, token, b)
}

func (c Client) SendMessage(msg *models.Message, token, receiver string) error {
	if msg == nil {
		return fmt.Errorf("message input is missing")
	}

	m, b := newMessage(minAPIVersion, nil, receiver, "", msg.ContentType.String(), msg.Text), new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(m); err != nil {
		return fmt.Errorf("encoding message")
	}
	log.Printf("making request with data %v", m)
	return c.makeRequest(sendMessageURL, token, b)
}
