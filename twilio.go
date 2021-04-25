package twilio

import (
	"net/url"

	"github.com/kevinburke/rest"
	"github.com/kevinburke/twilio-go"
)

type Client struct {
	*twilio.Client
	config *Config
}

type Config struct {
	AccountSID  string
	AuthToken   string
	SenderPhone string
}

func New(config *Config) *Client {
	return &Client{
		Client: twilio.NewClient(config.AccountSID, config.AuthToken, nil),
		config: config,
	}
}

func (client *Client) SendSMS(to string, message string, mediaURLs ...url.URL) error {
	var urls []*url.URL
	for _, url := range mediaURLs {
		urls = append(urls, &url)
	}

	_, err := client.Messages.SendMessage(client.config.SenderPhone, to, message, urls)
	if err != nil {
		switch v := err.(type) {
		case *rest.Error:
			return &Error{
				Detail:     v.Detail,
				Message:    v.Title,
				StatusCode: v.Status,
				Type:       v.Type,
			}
		}

		return err
	}

	return nil

}
