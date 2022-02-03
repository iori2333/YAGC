package cmd

import (
	"github.com/spf13/cobra"
	"yagc/db"
	"yagc/models"
)

func handleWriteTree(*cobra.Command, []string) {
	index := db.ReadIndex()
	tree := models.TreeObject{}
	tree.Parse(index)
	sha1, content := tree.GetSha1()
	db.WriteObject(sha1, content)
}

func getWriteTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "write-tree",
		Short:  "Write a tree object from the index",
		Run:    handleWriteTree,
		Hidden: true,
	}
	return cmd
}
