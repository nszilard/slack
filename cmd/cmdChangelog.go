package cmd

import (
	"fmt"
	"strings"

	"github.com/nszilard/slack/internal/util"
	"github.com/nszilard/slack/models"
	"github.com/nszilard/slack/pkg/templates"

	"github.com/spf13/cobra"
)

//----------------------------------------
// Variables
//----------------------------------------
var (
	changelogService string // Service name (how it is registered in Git)
	changelogVersion string // The new tag (semantic version)
	changelogCommits string // Commits to display
)

//----------------------------------------
// Cobra command
//----------------------------------------
var changelogCommand = &cobra.Command{
	Use:   "changelog",
	Short: "Sends a message using the Changelog template.",
	Long:  "Sends a changelog message using the provided values.",

	SilenceUsage:  true,
	SilenceErrors: true,

	PreRunE: checkChangelog,
	RunE:    sendChangelogMessage,
}

func sendChangelogMessage(cmd *cobra.Command, args []string) error {
	return templates.Changelog(client, changelogService, changelogVersion, changelogCommits)
}

//----------------------------------------
// Cobra command init
//----------------------------------------
func init() {
	// Local flags
	changelogCommand.Flags().StringVarP(&changelogService, models.ArgChangelogServiceFlag, "s", "", "Service name as it is registered in Git. (Required)")
	changelogCommand.Flags().StringVarP(&changelogVersion, models.ArgChangelogVersionFlag, "v", "", "The new tag. (Required)")

	// Fallback to Environment variables
	changelogService = util.FallbackEnvString(changelogService, models.ArgChangelogServiceEnv)
	changelogVersion = util.FallbackEnvString(changelogVersion, models.ArgChangelogVersionEnv)

	// Adding subcommand to the root command
	mainCmd.AddCommand(changelogCommand)
}

func checkChangelog(cmd *cobra.Command, args []string) error {
	changes := strings.Join(args, "\n")
	changelogCommits = util.FallbackEnvString(changes, models.ArgChangelogCommitsEnv)

	// Ensure required values are provided one way or another
	switch {
	case changelogService == "":
		return fmt.Errorf("missing required value: changelog: service")
	case changelogVersion == "":
		return fmt.Errorf("missing required value: changelog: version")
	case changelogCommits == "":
		return fmt.Errorf("missing required value: changelog: need to provide commits either as an argument or an environment variable")
	}

	return nil
}
