package slack

import "fmt"

// MessageBlockType defines the type for a Block
type MessageBlockType string

// MBTxxxxx defines the message header types
const (
	MBTHeader  MessageBlockType = "header"
	MBTDivider MessageBlockType = "divider"
	MBTSection MessageBlockType = "section"
	MBTAction  MessageBlockType = "actions"
	// Not supported by the app currently
	MBTImage    MessageBlockType = "image"
	MBTContext  MessageBlockType = "context"
	MBTFile     MessageBlockType = "file"
	MBTInput    MessageBlockType = "input"
	MBTRichText MessageBlockType = "rich_text"
)

// Block defines a block for a Slack message
// @See: https://api.slack.com/reference/block-kit/blocks
type Block struct {
	// Type is required in all Block types
	Type MessageBlockType `json:"type"`

	// These fields might be required depending on Block type
	Text     *TextObject      `json:"text,omitempty"`
	Fields   []*TextObject    `json:"fields,omitempty"`
	Elements []*ElementObject `json:"elements,omitempty"`
	BlockID  string           `json:"block_id,omitempty"`
}

func newSectionBlock(blockType MessageBlockType, text string, textType TextObjectType, fields ...*TextObject) *Block {
	block := &Block{
		Type: blockType,
	}

	if text != "" {
		block.Text = newTextBlockObject(textType, text, false, false)
	}

	if len(fields) > 0 {
		block.Fields = append(block.Fields, fields...)
	}

	return block
}

// NewHeaderBlock returns a new header block
func NewHeaderBlock(text string) *Block {
	return newSectionBlock(MBTHeader, text, TOTPlaintext)
}

// NewSectionBlock returns a new section block
func NewSectionBlock(text string, fields ...*TextObject) *Block {
	return newSectionBlock(MBTSection, text, TOTPlaintext, fields...)
}

// NewCodeSectionBlock returns a new section block which has the text in a code block
func NewCodeSectionBlock(text string) *Block {
	return newSectionBlock(MBTSection, fmt.Sprintf("```%v```", text), TOTMarkdown)
}

func newActionBlock(blockType MessageBlockType, blockID string, elements ...*ElementObject) *Block {
	block := &Block{
		Type:    blockType,
		BlockID: blockID,
	}

	if len(elements) > 0 {
		block.Elements = append(block.Elements, elements...)
	}

	return block
}

// NewActionBlock returns a new instance of an Action Block
func NewActionBlock(blockID string, elements ...*ElementObject) *Block {
	return newActionBlock(MBTAction, blockID, elements...)
}

// NewDividerBlock returns a new instance of a divider block
func NewDividerBlock() *Block {
	return newSectionBlock(MBTDivider, "", TOTPlaintext)
}
