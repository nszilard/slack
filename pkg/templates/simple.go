package templates

import (
	"github.com/nszilard/slack/internal/slack"
	"github.com/nszilard/slack/models"
)

// Simple will send a simple message
func Simple(client *models.Client, text string) error {
	msg := slack.NewBlockMessage(
		"You have received a new message",
		slack.NewSectionBlock(text),
	)
	client.SetMessage(&msg)

	return client.Send()
}
