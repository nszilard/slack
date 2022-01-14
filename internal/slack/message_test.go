package slack

import "testing"

func TestNewBlockMessage(t *testing.T) {
	actual := NewBlockMessage("notification message")

	if actual.Text != "notification message" {
		t.Errorf("new block message: unexpected text: %q (expected: %q)", actual.Text, "notification message")
	}

	if len(actual.Blocks) != 0 {
		t.Fatalf("new block message: unexpected number of blocks: %v (expected: %v)", len(actual.Blocks), 0)
	}
}

func TestMessage_AddBlock(t *testing.T) {
	actual := NewBlockMessage("notification message")
	actual.AddBlock(NewHeaderBlock("test"))

	if len(actual.Blocks) != 1 {
		t.Fatalf("add block: unexpected number of blocks: %v (expected: %v)", len(actual.Blocks), 1)
	}

	if actual.Blocks[0].Text.Text != "test" {
		t.Errorf("add block: unexpected text for new block: %q (expected: %q)", actual.Blocks[0].Text.Text, "test")
	}
}
