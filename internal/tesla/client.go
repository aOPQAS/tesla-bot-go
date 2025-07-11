package tesla

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	AccessToken string
	HttpClient  *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		AccessToken: token,
		HttpClient:  &http.Client{},
	}
}

func (c *Client) request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("Tesla API error: %s", resp.Status)
	}

	return resp, nil
}
