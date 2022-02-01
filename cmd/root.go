package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yagc",
	Short: "YAGC: Yet another git cli",
	Long:  "YAGC is a git cli that is built on top of go.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	subcommands := []*cobra.Command{
		getInitCmd(),
	}

	for _, subcommand := range subcommands {
		rootCmd.AddCommand(subcommand)
	}
}
