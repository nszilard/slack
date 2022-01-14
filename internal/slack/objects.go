package slack

import "fmt"

// TextObjectType defines the type for a TextObject
type TextObjectType string

// TOTxxxxx defines the TextObject types
const (
	TOTPlaintext TextObjectType = "plain_text"
	TOTMarkdown  TextObjectType = "mrkdwn"
)

// TextObject is a composition object that is used in most Text attribute of a Block object
// @See: https://api.slack.com/reference/block-kit/composition-objects#text
type TextObject struct {
	// Required fields
	Type TextObjectType `json:"type"` // The formatting to use for this text object. Can be one of plain_text or mrkdwn.
	Text string         `json:"text"` // The text for the block. This field accepts any of the standard text formatting markup when type is mrkdwn.

	// Optional fields
	Emoji    bool `json:"emoji,omitempty"`    // Indicates whether emojis in a text field should be escaped into the colon emoji format. This field is only usable when the type is plain text
	Verbatim bool `json:"verbatim,omitempty"` // Bool whether URLs will be auto-converted into links, conversation names will be link-ified, and certain mentions will be automatically parsed
}

func newTextBlockObject(elementType TextObjectType, text string, emoji, verbatim bool) *TextObject {
	obj := &TextObject{
		Type:     elementType,
		Text:     text,
		Verbatim: verbatim,
	}

	// Emoji field is only valid for "plain_text" type
	switch elementType {
	case TOTPlaintext:
		obj.Emoji = emoji
	}

	return obj
}

// NewTextObject returns a new TextObject. Setting bold to true will format the text as such.
func NewTextObject(text string, bold bool) *TextObject {
	switch bold {
	case true:
		return newTextBlockObject(TOTMarkdown, fmt.Sprintf("*%v*", text), false, false)

	default:
		return newTextBlockObject(TOTMarkdown, text, false, false)
	}
}

// NewKeyValueTextObject returns a new TextObject.
// The title has a bold format and the value is started in a new line.
func NewKeyValueTextObject(title, value string) *TextObject {
	return newTextBlockObject(TOTMarkdown, fmt.Sprintf("*%v*\n%v", title, value), false, false)
}

// NewCodeblockTextObject returns a new TextObject which has the content in a code block.
func NewCodeblockTextObject(text string) *TextObject {
	return newTextBlockObject(TOTMarkdown, fmt.Sprintf("```%v```", text), false, false)
}

// ElementObjectType defines the type for an ElementObject
type ElementObjectType string

// EOTxxxxx defines the ElementObject types
const (
	EOTButton ElementObjectType = "button"
	EOTImage  ElementObjectType = "image"
)

// ElementObject can be used inside Block objects
// @See: https://api.slack.com/reference/block-kit/block-elements
type ElementObject struct {
	// Required fields for all types
	Type ElementObjectType `json:"type"` // The type of element.

	// Shared by various element types
	Text     *TextObject `json:"text,omitempty"`      // A text object that defines the button's text. Can only be of type: `plain_text` or `text`. Max length: 75 chars.
	ActionID string      `json:"action_id,omitempty"` // Should be unique among all other action_ids in the containing block. Maximum length for this field is 255 characters.

	// Additional Button fields
	URL   string `json:"url,omitempty"`   // A URL to load in the user's browser when the button is clicked. Maximum length for this field is 3000 characters. (Optional)
	Value string `json:"value,omitempty"` // The value to send along with the interaction payload. Maximum length for this field is 2000 characters. (Optional)
	Style string `json:"style,omitempty"` // Decorates buttons with alternative visual color schemes. If provided, it can be either 'primary' or 'danger'. (Optional)

	// Additional Image fields
	ImageURL     string `json:"image_url,omitempty"` // The URL of the image to be displayed. (Required)
	ImageAltText string `json:"alt_text,omitempty"`  // A plain-text summary of the image. This should not contain any markup. (Required)
}

func newButtonElement(text, actionID, url, style string) *ElementObject {
	textObject := newTextBlockObject(TOTPlaintext, text, false, false)

	return &ElementObject{
		Type:     EOTButton,
		Text:     textObject,
		ActionID: actionID,
		URL:      url,
		Style:    style,
	}
}

// NewPrimaryButtonElement returns a new Button element which has the style: 'primary'
func NewPrimaryButtonElement(text, actionID, url string) *ElementObject {
	return newButtonElement(text, actionID, url, "primary")
}
