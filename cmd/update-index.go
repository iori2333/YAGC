package cmd

import (
	"github.com/spf13/cobra"
	"yagc/db"
)

func handleUpdateIndex(cmd *cobra.Command, args []string) {
	if cmd.Flag("add").Changed {
		db.CreateIndex(args)
	}
}

func getUpdateIndexCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "update-index",
		Short:  "Update the index",
		Long:   `Update the index`,
		Run:    handleUpdateIndex,
		Hidden: true,
	}

	cmd.Flags().Bool("add", false, "Add the objects to the index")
	cmd.Flags().Bool("remove", false, "Remove the objects from the index")
	cmd.Flags().Bool("replace", false, "Replace the index")
	cmd.Flags().Bool("refresh", false, "Refresh the index")

	cmd.Flags().BoolP("quiet", "q", false, "Suppress output")

	return cmd
}
