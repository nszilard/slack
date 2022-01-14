## slack changelog

Sends a message using the Changelog template.

### Synopsis

Sends a changelog message using the provided values.

```
slack changelog [flags]
```

### Options

```
  -s, --changelog-service string   Service name as it is registered in Git. (Required)
  -v, --changelog-version string   The new tag. (Required)
  -h, --help                       help for changelog
```

### Options inherited from parent commands

```
  -d, --debug                        Enables Debug logging.
  -C, --slack-channel string         Slack Channel where to send the message. (Optional)
  -O, --slack-org-id string          Slack Organization ID in the form of: Txxxxxx. (Required)
  -U, --slack-user string            Name to use when sending the message. (Optional)
  -I, --slack-user-image string      Link to an image to use as a profile icon when sending the message (Optional)
  -W, --slack-webhook-id string      Slack Webhook ID in the form of: Bxxxxxx. (Required)
  -T, --slack-webhook-token string   Slack Webhook token in the form of: xxxxxxx. (Required)
```

### SEE ALSO

* [slack](slack.md)	 - CLI to send Slack messages programmatically.

