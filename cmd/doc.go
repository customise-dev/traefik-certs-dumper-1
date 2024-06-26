package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docCmd represents the doc command.
var docCmd = &cobra.Command{
	Use:    "doc",
	Short:  "Generate documentation",
	Hidden: true,
	RunE: func(_ *cobra.Command, _ []string) error {
		return doc.GenMarkdownTree(rootCmd, "./docs")
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}
