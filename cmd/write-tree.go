package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"yagc/db"
)

func handleWriteTree(*cobra.Command, []string) {
	tree := db.ReadIndex()
	sha1, content := tree.GetSha1()
	db.WriteObject(sha1, content)
	log.Println(sha1)
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
