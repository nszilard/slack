package cmd

import (
	"strings"

	"github.com/nszilard/slack/pkg/templates"

	"github.com/spf13/cobra"
)

//----------------------------------------
// Cobra command
//----------------------------------------
var simpleCommand = &cobra.Command{
	Use:     "simple",
	Aliases: []string{"message", "text"},
	Short:   "Sends a simple text message.",
	Long:    "Sends a simple text message.",

	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: sendSimpleMessage,
}

func sendSimpleMessage(cmd *cobra.Command, args []string) error {
	text := strings.Join(args, "\n")
	return templates.Simple(client, text)
}

//----------------------------------------
// Cobra command init
//----------------------------------------
func init() {
	// Adding subcommand to the root command
	mainCmd.AddCommand(simpleCommand)
}
