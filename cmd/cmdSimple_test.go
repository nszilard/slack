package cmd

import (
	"encoding/json"
	"testing"

	"github.com/nszilard/slack/internal/slack"
)

func Test_sendSimpleMessage(t *testing.T) {
	args := []string{"test message"}

	err := sendSimpleMessage(simpleCommand, args)
	if err != nil {
		t.Errorf("send simple message: unexpected error: %v", err)
	}

	actual := slack.Message{}
	json.Unmarshal(client.HTTPClient.(*mockHTTPClient).pl, &actual)

	if actual.Blocks[0].Text.Text != "test message" {
		t.Errorf("send simple message: unexpected message: %q (expected: %q)", actual.Blocks[0].Text.Text, "test message")
	}
}
