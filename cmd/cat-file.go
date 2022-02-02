package cmd

import (
	"log"
	"os"
	"yagc/db"
	"yagc/util"

	"github.com/spf13/cobra"
)

func handleCatFile(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Usage()
		return
	}

	var object, objType string

	useType, useSize, useError, usePretty :=
		cmd.Flag("type").Changed,
		cmd.Flag("size").Changed,
		cmd.Flag("error").Changed,
		cmd.Flag("pretty").Changed

	if len(args) == 1 {
		if !useType && !useSize && !useError {
			log.Fatalln("Type is required")
		}
		object = args[0]
	} else {
		objType = args[0]
		object = args[1]
	}

	content := db.Find(object)
	realType, size, content, err := util.DecodeObject(content)

	if useError {
		if err != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}

	if err != nil {
		log.Fatalf("Failed to find or decode object %s: %s\n", object, err)
	}

	if useType {
		log.Println(realType)
	} else if useSize {
		log.Println(size)
	} else if usePretty {
		switch realType {
		case "blob":
			log.Println(string(content))
		case "tree":
			log.Println(string(content))
		case "commit":
			log.Println(string(content))
		}
	} else if realType == objType {
		log.Println(string(content))
	} else {
		log.Fatalf("Object %s is not a %s\n", object, objType)
	}
}

func getCatFileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cat-file [type] object",
		Short: "Provide content or type and size information for repository objects",
		Long:  `The command provides the content or the type of an object in the repository. The type is required unless -t or -p is used to find the object type, or -s is used to find the object size.`,
		Run:   handleCatFile,
	}

	cmd.Flags().BoolP("type", "t", false, "Print the type of the object")

	cmd.Flags().BoolP("pretty", "p", false, "Pretty print the object")

	cmd.Flags().BoolP("size", "s", false, "Print the size of the object")

	cmd.Flags().BoolP("error", "e", false, "Test whether an object exists")

	return cmd
}
