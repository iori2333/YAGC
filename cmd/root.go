package cmd

import (
	"log"
	"os"
	"yagc/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yagc",
	Short: "YAGC: Yet another git cli",
	Long:  "YAGC is a git cli that is built on top of go.",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("version").Value.String() == "true" {
			log.Printf("yagc version %s\n", config.Version)
		} else {
			_ = cmd.Help()
		}
	},
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
		getHashObjectCmd(),
		getCatFileCmd(),
		getWriteTreeCmd(),
		getUpdateIndexCmd(),
		getCommitCmd(),
	}

	for _, subcommand := range subcommands {
		rootCmd.AddCommand(subcommand)
	}

	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
}
