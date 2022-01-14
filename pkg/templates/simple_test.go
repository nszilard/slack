package templates

import "testing"

func TestSimple(t *testing.T) {
	expected := "simple text message"

	err := Simple(mockClient, expected)
	if err != nil {
		t.Errorf("simple message: something went wrong: %v", err)
	}

	actual := mockClient.HTTPClient.(*mockHTTPClientSuccess).sent
	if actual.Blocks[0].Text.Text != expected {
		t.Errorf("simple message: unexpected data: %q (expected: %q)", actual.Blocks[0].Text.Text, expected)
	}
}
