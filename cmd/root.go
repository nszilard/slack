package cmd

import (
	"fmt"

	"github.com/nszilard/slack/internal/util"
	"github.com/nszilard/slack/models"

	"github.com/spf13/cobra"

	"github.com/nszilard/log"
)

//----------------------------------------
// Variables (including shared variables)
//----------------------------------------
var (
	slackOrgID     string         // Slack Organization ID (Txxxxxx)
	slackAppID     string         // Slack Webhook ID (Bxxxxxx)
	slackToken     string         // Slack Webhook Token (xxxxxxx)
	slackChannel   string         // Slack channel to send the message
	slackUser      string         // Display name to send the message as
	slackUserImage string         // Image to use when sending the message
	debug          bool           // Show debug logs
	client         *models.Client // Slack client
)

//----------------------------------------
// Cobra command
//----------------------------------------
var mainCmd = &cobra.Command{
	Use:   "slack",
	Short: "CLI to send Slack messages programmatically.",
	Long:  "A simple CLI tool to send Slack messages programmatically using pre-defined templates.",

	PersistentPreRunE: checkAndSet,
	DisableAutoGenTag: true,
	SilenceUsage:      true,
	SilenceErrors:     true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the mainCmd.
func Execute() error {
	return mainCmd.Execute()
}

//----------------------------------------
// Cobra command init
//----------------------------------------
func init() {
	// Global persistent flags that are required to have a value
	mainCmd.PersistentFlags().StringVarP(&slackOrgID, models.ArgSlackOrgIDFlag, "O", "", "Slack Organization ID in the form of: Txxxxxx. (Required)")
	mainCmd.PersistentFlags().StringVarP(&slackAppID, models.ArgSlackWebhookIDFlag, "W", "", "Slack Webhook ID in the form of: Bxxxxxx. (Required)")
	mainCmd.PersistentFlags().StringVarP(&slackToken, models.ArgSlackWebhookTokenFlag, "T", "", "Slack Webhook token in the form of: xxxxxxx. (Required)")

	// Global persistent flags that are not required to be set
	mainCmd.PersistentFlags().StringVarP(&slackChannel, models.ArgSlackChannelFlag, "C", "", "Slack Channel where to send the message. (Optional)")
	mainCmd.PersistentFlags().StringVarP(&slackUser, models.ArgSlackUsernameFlag, "U", "", "Name to use when sending the message. (Optional)")
	mainCmd.PersistentFlags().StringVarP(&slackUserImage, models.ArgSlackUserImageFlag, "I", "", "Link to an image to use as a profile icon when sending the message (Optional)")
	mainCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enables Debug logging.")

	// Fallback to Environment variables
	slackOrgID = util.FallbackEnvString(slackOrgID, models.ArgSlackOrgIDEnv)
	slackAppID = util.FallbackEnvString(slackAppID, models.ArgSlackWebhookIDEnv)
	slackToken = util.FallbackEnvString(slackToken, models.ArgSlackWebhookTokenEnv)
	slackChannel = util.FallbackEnvString(slackChannel, models.ArgSlackChannelEnv)
	slackUser = util.FallbackEnvString(slackUser, models.ArgSlackUsernameEnv)
	slackUserImage = util.FallbackEnvString(slackUserImage, models.ArgSlackUserImageEnv)
}

func checkAndSet(cmd *cobra.Command, args []string) error {
	// Ensure required values are provided one way or another
	switch {
	case slackOrgID == "":
		return fmt.Errorf("missing required values: slack: Organization ID")
	case slackAppID == "":
		return fmt.Errorf("missing required values: slack: Webhook ID")
	case slackToken == "":
		return fmt.Errorf("missing required values: slack: Webhook token")
	}

	// Set log level
	log.SetLevelInfo()
	if debug {
		log.SetLevelDebug()
	}

	// Initiate client
	client = models.NewClient(slackOrgID, slackAppID, slackToken).WithChannel(slackChannel).WithUser(slackUser).WithUserImage(slackUserImage)

	return nil
}
