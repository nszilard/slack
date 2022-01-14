package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsLocation string

//----------------------------------------
// Cobra command
//----------------------------------------
var documentationCmd = &cobra.Command{
	Use:     "docs",
	Aliases: []string{"doc", "documentation"},
	Short:   "Generates markdown documentation",

	Hidden: true,

	Args: cobra.NoArgs,
	RunE: generateMarkdown,
}

func generateMarkdown(cmd *cobra.Command, args []string) error {
	return doc.GenMarkdownTree(cmd.Root(), docsLocation)
}

//----------------------------------------
// Cobra command init
//----------------------------------------
func init() {
	// Adding documentation command to the root command
	mainCmd.AddCommand(documentationCmd)

	docsLocation = "docs"
}
