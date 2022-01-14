package slack

// Message describes the payload that the Slack API can receive
// @See: https://api.slack.com/reference/messaging/payload
type Message struct {
	Text    string   `json:"text"`              // It is used as a fallback string to display in notifications
	Blocks  []*Block `json:"blocks"`            // An array of layout blocks.
	Channel string   `json:"channel,omitempty"` // Channel where the message should be published

	UserName  string `json:"username,omitempty"`   // To specify the username for the published message.
	IconURL   string `json:"icon_url,omitempty"`   // To specify a URL to an image to use as the profile photo alongside the message.
	IconEmoji string `json:"icon_emoji,omitempty"` // To specify an emoji (using colon shortcodes, eg. :wave:) to use as the profile photo alongside the message.

	MarkdownSupport bool `json:"mrkdwn,omitempty"` // Determines whether the text field is rendered according to mrkdwn formatting or not.
}

// AddBlock appends a block to the end of the existing list of blocks
func (msg *Message) AddBlock(block *Block) {
	msg.Blocks = append(msg.Blocks, block)
}

// NewMessage creates a new Message that contains one or more blocks to be displayed
func NewBlockMessage(notification string, blocks ...*Block) Message {
	return Message{
		Text:   notification,
		Blocks: blocks,
	}
}
