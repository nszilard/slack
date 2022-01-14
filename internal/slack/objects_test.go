package slack

import "testing"

func TestNewTextObject(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		bold     bool
		expected string
	}{
		{
			name:     "simple message",
			text:     "test",
			expected: "test",
		},
		{
			name:     "bold message",
			text:     "test",
			bold:     true,
			expected: "*test*",
		},
	}
	for _, tt := range tests {
		actual := NewTextObject(tt.text, tt.bold)

		if actual.Type != TOTMarkdown {
			t.Errorf("(case: %q) unexpected type: %q, expected: %q", tt.name, actual.Type, TOTMarkdown)
		}

		if actual.Text != tt.expected {
			t.Errorf("(case: %q) unexpected text: %q, expected: %q)", tt.name, actual.Text, tt.expected)
		}
	}
}

func TestNewKeyValueTextObject(t *testing.T) {
	actual := NewKeyValueTextObject("title", "value")

	if actual.Text != "*title*\nvalue" {
		t.Errorf("key value text object: unexpected text value")
	}
}

func TestNewCodeblockTextObject(t *testing.T) {
	actual := NewCodeblockTextObject("code")

	if actual.Text != "```code```" {
		t.Errorf("code block text object: unexpected text value")
	}
}

func TestNewPrimaryButtonElement(t *testing.T) {
	actual := NewPrimaryButtonElement("testButton", "testActionID", "testURL")

	if actual.Type != EOTButton {
		t.Errorf("action block: wrong block type: %q", actual.Type)
	}

	if actual.Text.Text != "testButton" {
		t.Fatalf("action block: unexpected text in button: %q", actual.Text.Text)
	}

	if actual.ActionID != "testActionID" {
		t.Fatalf("action block: unexpected actionID for button: %q", actual.ActionID)
	}

	if actual.URL != "testURL" {
		t.Fatalf("action block: unexpected URL for button: %q", actual.URL)
	}
}
