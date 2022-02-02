package cmd

import (
	"log"

	"yagc/db"
	"yagc/models"

	"github.com/spf13/cobra"
)

func handleHashObject(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln("missing argument: filename")
	}

	t := cmd.Flag("type").Value.String()
	if t != "blob" {
		log.Fatalln("unsupported object type:", t)
	}

	write := cmd.Flag("write").Value.String() == "true"

	obj := models.BlobObject{
		File: args[0],
	}

	sha1, content := obj.GetSha1()
	log.Println(sha1)

	if write {
		db.Write(sha1, content)
	}
}

func getHashObjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash-object",
		Short: "Compute object ID and optionally creates a blob from a file",
		Long:  `Compute object ID and optionally creates a blob from a file. The object ID is printed to standard output.`,
		Run:   handleHashObject,
	}

	cmd.Flags().StringP("type", "t", "blob", "Specify the type of the object")

	cmd.Flags().BoolP("write", "w", false, "Actually write the object into the database")

	cmd.Flags().Bool("stdin", false, "Read standard input instead of a file")

	return cmd
}
