package cmd

import (
	"log"
	"os"

	"yagc/db"
	"yagc/models"

	"github.com/spf13/cobra"
)

func handleHashObject(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln("Missing argument: file")
	}

	file := args[0]
	var obj models.Object

	switch t := cmd.Flag("type").Value.String(); t {
	case "blob":
		obj = &models.BlobObject{
			File: file,
		}
	case "tree":
		tree := models.TreeObject{}
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read file %s: %s\n", file, err)
		}
		obj = tree.Parse(content)
	case "commit":
		log.Fatalln("Not implemented")
	default:
		log.Fatalf("Invalid type received: %s\n", t)
	}

	sha1, content := obj.GetSha1()
	log.Println(sha1)

	if cmd.Flag("write").Changed {
		db.WriteObject(sha1, content)
	}
}

func getHashObjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash-object file",
		Short: "Compute object ID and optionally creates a blob from a file",
		Long: `Computes the object ID value for an object with specified type with the contents of the named file
(which can be outside of the work tree), and optionally writes the resulting object into the object database.
Reports its object ID to its standard output. When <type> is not specified, it defaults to "blob".`,
		Run:    handleHashObject,
		Hidden: true,
	}

	cmd.Flags().StringP("type", "t", "blob", "Specify the type of the object")
	cmd.Flags().BoolP("write", "w", false, "Actually write the object into the database")
	cmd.Flags().Bool("stdin", false, "Read standard input instead of a file")

	return cmd
}
