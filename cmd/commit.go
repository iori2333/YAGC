package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"yagc/db"
)

func handleCommit(cmd *cobra.Command, _ []string) {
	message := cmd.Flag("message").Value.String()
	if message == "" {
		log.Fatalf("Message is required to commit")
	}

	db.Commit(message)
}

func getCommitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "Commit changes to the repository",
		Long:  `Commit changes to the repository`,
		Run:   handleCommit,
	}

	cmd.Flags().StringP("message", "m", "", "Commit message")
	cmd.Flags().StringP("all", "a", "", "Commit all changes")

	return cmd
}
