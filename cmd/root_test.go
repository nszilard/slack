package cmd

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/nszilard/slack/models"

	"github.com/spf13/cobra"
)

type mockHTTPClient struct {
	pl []byte
}

func (m *mockHTTPClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	b, _ := io.ReadAll(body)
	m.pl = b

	return &http.Response{Status: "200 OK", StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("success"))}, nil
}

func init() {
	client = &models.Client{
		HTTPClient: &mockHTTPClient{},
	}
}

func TestExecute(t *testing.T) {
	err := Execute()
	if err != nil {
		t.Errorf("execute: unexpected error: %v", err)
	}
}

func Test_checkAndSet(t *testing.T) {
	tests := []struct {
		name         string
		orgID        string
		webhookID    string
		webhookToken string
		cmd          *cobra.Command
		args         []string
		debug        bool
		wantErr      bool
	}{
		{
			name:         "Should error for empty SlackOrgID",
			cmd:          mainCmd,
			orgID:        "",
			webhookID:    "test",
			webhookToken: "test",
			wantErr:      true,
		},
		{
			name:         "Should error for empty SlackAppID",
			cmd:          mainCmd,
			orgID:        "test",
			webhookID:    "",
			webhookToken: "test",
			wantErr:      true,
		},
		{
			name:         "Should error for empty SlackToken",
			cmd:          mainCmd,
			orgID:        "test",
			webhookID:    "test",
			webhookToken: "",
			wantErr:      true,
		},
		{
			name:         "Should set debug level",
			cmd:          mainCmd,
			orgID:        "test",
			webhookID:    "test",
			webhookToken: "test",
			debug:        true,
		},
	}
	for _, tt := range tests {
		slackOrgID = tt.orgID
		slackAppID = tt.webhookID
		slackToken = tt.webhookToken
		debug = tt.debug

		if err := checkAndSet(tt.cmd, tt.args); (err != nil) != tt.wantErr {
			t.Errorf("(case %q) error: %v, wanted error: %v", tt.name, err, tt.wantErr)
		}
	}
}
