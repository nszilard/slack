package templates

import (
	"fmt"
	"strings"
	"time"

	"github.com/nszilard/slack/internal/slack"
	"github.com/nszilard/slack/models"
)

// Changelog provides a template to send Changlog messages
func Changelog(client *models.Client, service, version, changes string) error {
	// Divider
	divider := slack.NewDividerBlock()

	// Header
	headerSection := slack.NewHeaderBlock(fmt.Sprintf("Changelog - %v", time.Now().Local().Format(time.RFC1123)))

	// Context
	serviceBlock := slack.NewKeyValueTextObject("Service:", service)
	versionBlock := slack.NewKeyValueTextObject("Version:", version)
	contextSection := slack.NewSectionBlock("", serviceBlock, versionBlock)

	// Commits
	commitsTitleSection := slack.NewSectionBlock("", slack.NewTextObject("Commits:", true))
	commitsSection := slack.NewCodeSectionBlock(strings.ReplaceAll(changes, "\t", ""))

	msg := slack.NewBlockMessage(
		fmt.Sprintf("A new version has been released for: %q", service),
		headerSection,
		divider,
		contextSection,
		commitsTitleSection,
		commitsSection,
		divider,
	)
	client.SetMessage(&msg)

	return client.Send()
}
