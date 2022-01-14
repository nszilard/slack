## slack

CLI to send Slack messages programmatically.

### Synopsis

A simple CLI tool to send Slack messages programmatically using pre-defined templates.

### Options

```
  -d, --debug                        Enables Debug logging.
  -h, --help                         help for slack
  -C, --slack-channel string         Slack Channel where to send the message. (Optional)
  -O, --slack-org-id string          Slack Organization ID in the form of: Txxxxxx. (Required)
  -U, --slack-user string            Name to use when sending the message. (Optional)
  -I, --slack-user-image string      Link to an image to use as a profile icon when sending the message (Optional)
  -W, --slack-webhook-id string      Slack Webhook ID in the form of: Bxxxxxx. (Required)
  -T, --slack-webhook-token string   Slack Webhook token in the form of: xxxxxxx. (Required)
```

### SEE ALSO

* [slack changelog](slack_changelog.md)	 - Sends a message using the Changelog template.
* [slack simple](slack_simple.md)	 - Sends a simple text message.

