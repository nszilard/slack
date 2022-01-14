package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nszilard/slack/internal/slack"

	"github.com/nszilard/log"
)

const (
	// apiURL of the Slack Hooks API
	apiURL = "https://hooks.slack.com/services"
)

// httpClient defines the minimal interface needed for an http.Client to be implemented.
type httpClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

// Client for the slack api.
type Client struct {
	api        string
	secret     string
	channel    string
	userName   string
	userImage  string
	msg        *slack.Message
	HTTPClient httpClient
}

// NewClient builds a slack client from the provided token and options.
func NewClient(orgID, webhookID, webhookToken string) *Client {
	s := &Client{
		api:        apiURL,
		secret:     fmt.Sprintf("%v/%v/%v", orgID, webhookID, webhookToken),
		msg:        nil,
		HTTPClient: &http.Client{},
	}

	return s
}

// WithAPI overwrites the API URL of the client
func (client *Client) WithAPI(url string) *Client {
	client.api = url
	return client
}

// WithClient overwrites the HTTP Client
func (client *Client) WithClient(c httpClient) *Client {
	client.HTTPClient = c
	return client
}

// WithChannel overwrites the Slack channel
func (client *Client) WithChannel(channel string) *Client {
	client.channel = channel
	return client
}

// WithUser overwrites the Display name when sending the message
func (client *Client) WithUser(name string) *Client {
	client.userName = name
	return client
}

// WithUserImage overwrites the User image link when sending the message
func (client *Client) WithUserImage(link string) *Client {
	client.userImage = link
	return client
}

// SetMessage sets the msg object on the client
func (client *Client) SetMessage(msg *slack.Message) {
	client.msg = msg
}

// Send is responsible for sending the Slack message
func (client *Client) Send() error {
	setOptionsOnMessage(client)

	pl, err := encodeMessage(client.msg)
	if err != nil {
		return fmt.Errorf("send: marshall payload: %v", err)
	}

	url := fmt.Sprintf("%v/%v", client.api, client.secret)
	resp, err := client.HTTPClient.Post(url, "application/json", bytes.NewBuffer(pl))
	if err != nil {
		return fmt.Errorf("send: execute request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return handleResponse(resp, body, err)
}

func setOptionsOnMessage(client *Client) {
	if client.channel != "" {
		client.msg.Channel = client.channel
	}
	if client.userName != "" {
		client.msg.UserName = client.userName
	}
	if client.userImage != "" {
		client.msg.IconURL = client.userImage
	}
}

func encodeMessage(data interface{}) ([]byte, error) {
	pl, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}

	// Output JSON message
	log.Debug(string(pl))

	return pl, nil
}

func handleResponse(resp *http.Response, body []byte, err error) error {
	switch resp.StatusCode {
	case http.StatusOK:
		log.Debug(string(body))
		return nil
	default:
		if err != nil {
			return fmt.Errorf("send: unexpected response: read response body: %v", err)
		}

		return fmt.Errorf("send: unexpected response: %v: %v", resp.Status, string(body))
	}
}
