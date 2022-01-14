package templates

import (
	"fmt"
	"testing"

	"github.com/nszilard/slack/internal/slack"
)

func TestChangelog(t *testing.T) {
	service := "test"
	version := "1.0.0"
	changes := "- feat: Some really useful stuff"

	expected := struct {
		numberOfBlocks int
		headerType     slack.MessageBlockType
		dividerType    slack.MessageBlockType
		service        string
		version        string
		commitTitle    string
		commitValue    string
		button         string
	}{
		numberOfBlocks: 6,
		headerType:     slack.MBTHeader,
		dividerType:    slack.MBTDivider,
		service:        fmt.Sprintf("*Service:*\n%v", service),
		version:        fmt.Sprintf("*Version:*\n%v", version),
		commitTitle:    "*Commits:*",
		commitValue:    fmt.Sprintf("```%v```", changes),
	}

	err := Changelog(mockClient, service, version, changes)
	if err != nil {
		t.Errorf("changelog: something went wrong: %v", err)
	}

	actual := mockClient.HTTPClient.(*mockHTTPClientSuccess).sent

	// Check if there are exactly as many blocks as we want
	if len(actual.Blocks) != expected.numberOfBlocks {
		t.Fatalf("changelog: expected %v, but got %v", expected.numberOfBlocks, len(actual.Blocks))
	}

	// Test all blocks
	if actual.Blocks[0].Type != expected.headerType {
		t.Errorf("changelog: expected the first block to be a header, but got: %v", actual.Blocks[0].Type)
	}

	if actual.Blocks[1].Type != expected.dividerType {
		t.Errorf("changelog: expected the second block to be a divider, but got: %v", actual.Blocks[1].Type)
	}

	if actual.Blocks[2].Fields[0].Text != expected.service {
		t.Errorf("changelog: unexpected service block")
	}

	if actual.Blocks[2].Fields[1].Text != expected.version {
		t.Errorf("changelog: unexpected version block")
	}

	if actual.Blocks[3].Fields[0].Text != expected.commitTitle {
		t.Errorf("changelog: unexpected commit title block")
	}

	if actual.Blocks[4].Text.Text != expected.commitValue {
		t.Errorf("changelog: unexpected commit value block")
	}

	if actual.Blocks[5].Type != expected.dividerType {
		t.Errorf("changelog: expected the sixth block to be a divider, but got: %v", actual.Blocks[5].Type)
	}
}
