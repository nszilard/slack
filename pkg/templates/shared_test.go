package templates

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/nszilard/slack/internal/slack"
	"github.com/nszilard/slack/models"
)

type mockHTTPClientSuccess struct {
	sent *slack.Message
}

func (m *mockHTTPClientSuccess) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	pl := &slack.Message{}
	json.Unmarshal(data, pl)
	m.sent = pl

	return &http.Response{Status: "200 OK", StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("success"))}, nil
}

var mockClient *models.Client

func init() {
	// Initiate a client to use in unit tests
	mockClient = models.NewClient("org", "app", "token").WithClient(&mockHTTPClientSuccess{})
}
