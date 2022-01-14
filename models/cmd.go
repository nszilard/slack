package models

// Root command environment variables and flags
const (
	ArgSlackOrgIDEnv        = "SLACK_ORG_ID"
	ArgSlackWebhookIDEnv    = "SLACK_WEBHOOK_ID"
	ArgSlackWebhookTokenEnv = "SLACK_WEBHOOK_TOKEN"
	ArgSlackChannelEnv      = "SLACK_CHANNEL"
	ArgSlackUsernameEnv     = "SLACK_USER"
	ArgSlackUserImageEnv    = "SLACK_USER_IMAGE"

	ArgSlackOrgIDFlag        = "slack-org-id"
	ArgSlackWebhookIDFlag    = "slack-webhook-id"
	ArgSlackWebhookTokenFlag = "slack-webhook-token"
	ArgSlackChannelFlag      = "slack-channel"
	ArgSlackUsernameFlag     = "slack-user"
	ArgSlackUserImageFlag    = "slack-user-image"
)

// Changelog command environment variables and flags
const (
	ArgChangelogGroupEnv   = "CHANGELOG_GROUP"
	ArgChangelogServiceEnv = "CHANGELOG_SERVICE"
	ArgChangelogVersionEnv = "CHANGELOG_VERSION"
	ArgChangelogCommitsEnv = "CHANGELOG_COMMITS"

	ArgChangelogGroupFlag   = "changelog-group"
	ArgChangelogServiceFlag = "changelog-service"
	ArgChangelogVersionFlag = "changelog-version"
)
