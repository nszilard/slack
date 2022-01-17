package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func Test_generateMarkdown(t *testing.T) {
	folder := "temp"
	err := os.Mkdir(folder, 0755)
	if err != nil {
		t.Fatalf("failed to create temp folder")
	}
	docsLocation = folder
	defer os.RemoveAll(folder)

	err = generateMarkdown(&cobra.Command{Use: "test", DisableAutoGenTag: true}, nil)
	if err != nil {
		t.Errorf("docs: unexpected error: %v", err)
	}

	_, err = os.ReadFile(fmt.Sprintf("%v/%v", folder, "test.md"))
	if err != nil {
		t.Errorf("docs: unexpected error reading generated md file: %v", err)
	}
}
