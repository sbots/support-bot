package viber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"support-bot/models"
)

const apiEndpoint = "https://chatapi.viber.com/pa/set_webhook"

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

func (c Client) ConnectNewBot(token, path string) error {
	sub, b := &models.ViberSubscriptionRequest{
		Url:        path,
		EventTypes: models.AllViberEventTypes,
		SendName:   true,
		SendPhoto:  true,
	}, new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(sub); err != nil {
		return fmt.Errorf("encoding subsc")
	}
	req, err := http.NewRequest(http.MethodPost, apiEndpoint, b)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("X-Viber-Auth-Token", token)

	rsp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("doing http request: %w", err)
	}
	defer rsp.Body.Close()

	return handleResponse(rsp)
}

func handleResponse(rsp *http.Response) error {
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	var response models.ViberSubscriptionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}

	if response.Status == 0 {
		return nil
	}

	if msg, ok := errorCodes[response.Status]; ok {
		return fmt.Errorf("viber server respond with error: %s", msg)
	}
	return fmt.Errorf("viber server error")
}
