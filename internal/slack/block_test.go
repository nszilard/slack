package slack

import "testing"

func TestNewHeaderBlock(t *testing.T) {
	actual := NewHeaderBlock("test")

	if actual.Type != MBTHeader {
		t.Errorf("header block: unexpected block type: %q (expected: %q)", actual.Type, MBTHeader)
	}

	if actual.Text.Text != "test" {
		t.Errorf("header block: unexpected text: %q (expected: %q)", actual.Text.Text, "test")
	}
}

func TestNewSectionBlock(t *testing.T) {
	actual := NewSectionBlock("", NewTextObject("field test", false))

	if actual.Type != MBTSection {
		t.Errorf("section block: unexpected block type: %q (expected: %q)", actual.Type, MBTSection)
	}

	if len(actual.Fields) != 1 {
		t.Fatalf("section block: unexpected number of fields: %v (expected: %v)", len(actual.Fields), 1)
	}

	if actual.Fields[0].Text != "field test" {
		t.Errorf("section block: unexpected text for field: %q (expected: %q)", actual.Fields[0].Text, "field test")
	}
}

func TestNewCodeSectionBlock(t *testing.T) {
	actual := NewCodeSectionBlock("inside a code block")

	if actual.Type != MBTSection {
		t.Errorf("code section block: unexpected block type: %q (expected: %q)", actual.Type, MBTSection)
	}

	if len(actual.Fields) != 0 {
		t.Fatalf("code section block: unexpected number of fields: %v (expected: %v)", len(actual.Fields), 0)
	}

	if actual.Text.Text != "```inside a code block```" {
		t.Errorf("code section block: unexpected text: %q (expected: %q)", actual.Text.Text, "```inside a code block```")
	}
}

func TestNewActionBlock(t *testing.T) {
	actual := NewActionBlock("testBlockID", NewPrimaryButtonElement("testButton", "testActionID", "testURL"))

	if actual.Type != MBTAction {
		t.Errorf("action block: unexpected block type: %q (expected: %q)", actual.Type, MBTAction)
	}

	if actual.BlockID != "testBlockID" {
		t.Errorf("action block: unexpected blockDI: %q (expected: %q)", actual.BlockID, "testBlockID")
	}

	if len(actual.Elements) != 1 {
		t.Fatalf("action block: unexpected number of elements: %v (expected: %v)", len(actual.Elements), 1)
	}

	if actual.Elements[0].Text.Text != "testButton" {
		t.Errorf("action block: unexpected text in button: %q (expected: %q)", actual.Elements[0].Text.Text, "testButton")
	}

	if actual.Elements[0].ActionID != "testActionID" {
		t.Errorf("action block: unexpected actionID for button: %q (expected: %q)", actual.Elements[0].ActionID, "testActionID")
	}

	if actual.Elements[0].URL != "testURL" {
		t.Errorf("action block: unexpected URL for button: %q (expected: %q)", actual.Elements[0].URL, "testURL")
	}
}

func TestNewDividerBlock(t *testing.T) {
	actual := NewDividerBlock()

	if actual.Type != MBTDivider {
		t.Errorf("divider block: unexpected block type: %q (expected: %q)", actual.Type, MBTDivider)
	}
}
