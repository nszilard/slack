package cmd

import (
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

	if err := generateMarkdown(&cobra.Command{Use: "test"}, nil); err != nil {
		t.Errorf("docs: unexpected error: %v", err)
	}
}
