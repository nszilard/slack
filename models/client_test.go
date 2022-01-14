package models

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/nszilard/slack/internal/slack"
)

type mockHTTPClient struct {
	called bool
}

func (m *mockHTTPClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	m.called = true
	return &http.Response{Status: "200 OK", StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("success"))}, nil
}

type mockHTTPClientError struct{}

func (m *mockHTTPClientError) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return nil, fmt.Errorf("mock error")
}

type mockHTTPClientErrorResponse struct{}

func (m *mockHTTPClientErrorResponse) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return &http.Response{Status: "400 Bad Request", StatusCode: http.StatusBadRequest, Body: io.NopCloser(strings.NewReader("invalid blocks"))}, nil
}

func TestNewClient(t *testing.T) {
	actual := NewClient("orgID", "webhookID", "webhookToken")

	if actual.api != apiURL {
		t.Errorf("unexpected api url for new client: %q (expected: %q)", actual.api, apiURL)
	}

	expectedSecret := "orgID/webhookID/webhookToken"
	if actual.secret != expectedSecret {
		t.Errorf("unexpected secret for new client: %q (expected: %q)", actual.secret, expectedSecret)
	}
}

func TestClient_WithAPI(t *testing.T) {
	url := "mockAPI"
	actual := NewClient("orgID", "webhookID", "webhookToken").WithAPI(url)

	if actual.api != url {
		t.Errorf("unexpected api url: %q (expected: %q)", actual.api, url)
	}
}

func TestClient_WithClient(t *testing.T) {
	mockClient := &mockHTTPClient{}
	actual := NewClient("orgID", "webhookID", "webhookToken").WithClient(mockClient)
	actual.HTTPClient.Post("url", "test", nil)

	if mockClient.called != true {
		t.Errorf("unexpected client: WithClient failed to replace the HTTP client")
	}
}

func TestClient_WithChannel(t *testing.T) {
	channel := "test"
	actual := NewClient("orgID", "webhookID", "webhookToken").WithChannel(channel)

	if actual.channel != channel {
		t.Errorf("unexpected channel: %q (expected: %q)", actual.channel, channel)
	}
}

func TestClient_WithUser(t *testing.T) {
	user := "test"
	actual := NewClient("orgID", "webhookID", "webhookToken").WithUser(user)

	if actual.userName != user {
		t.Errorf("unexpected userName: %q (expected: %q)", actual.userName, user)
	}
}

func TestClient_WithUserImage(t *testing.T) {
	link := "test"
	actual := NewClient("orgID", "webhookID", "webhookToken").WithUserImage(link)

	if actual.userImage != link {
		t.Errorf("unexpected userImage: %q (expected: %q)", actual.channel, link)
	}
}

func TestClient_SetMessage(t *testing.T) {
	msg := &slack.Message{Text: "test"}
	actual := NewClient("orgID", "webhookID", "webhookToken")
	actual.SetMessage(msg)

	if actual.msg.Text != "test" {
		t.Errorf("failed to set message on client")
	}
}

func TestClient_Send(t *testing.T) {
	tests := []struct {
		name    string
		client  httpClient
		msg     *slack.Message
		wantErr bool
	}{
		{
			name:   "Success",
			client: &mockHTTPClient{},
			msg:    &slack.Message{Text: "test"},
		},
		{
			name:    "Fail: error when sending message",
			client:  &mockHTTPClientError{},
			msg:     &slack.Message{Text: "test"},
			wantErr: true,
		},
		{
			name:    "Fail: response is not 200",
			client:  &mockHTTPClientErrorResponse{},
			msg:     &slack.Message{Text: "test"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient("a", "b", "c").WithChannel("test").WithUser("test").WithUserImage("test").WithClient(tt.client)
			client.SetMessage(tt.msg)
			if err := client.Send(); (err != nil) != tt.wantErr {
				t.Errorf("Case %q: error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func Test_encodeMessageError(t *testing.T) {
	x := map[string]interface{}{
		"foo": make(chan int),
	}

	_, err := encodeMessage(x)
	if err == nil {
		t.Errorf("expected error, recevied none")
	}
}

func Test_handleResponseError(t *testing.T) {
	err := handleResponse(&http.Response{Status: "400 Bad Request", StatusCode: http.StatusBadRequest, Body: io.NopCloser(strings.NewReader("invalid blocks"))}, []byte("mock"), fmt.Errorf("unable to read body"))
	if err == nil {
		t.Errorf("expected error, recevied none")
	}
}
