package viber

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (c Client) makeRequest(url, token string, b io.Reader) error {
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("X-Viber-Auth-Token", token)

	rsp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("doing http request: %w", err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	var response SubscriptionResponse
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
