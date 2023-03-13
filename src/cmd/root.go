package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "easypdf",
	Short:   "Create pdfs from markdown with a single command",
	Long:    "Create pdfs from markdown with a single command",
	Version: "v0.0.1",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
