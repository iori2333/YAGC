package cmd

import (
	"log"
	"os"
	"path"
	"yagc/config"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const FileMode = 0644

func handleInit(cmd *cobra.Command, args []string) {
	var filePath string

	isBare := cmd.Flag("bare").Value.String() == "true"

	if len(args) == 0 {
		filePath, _ = os.Getwd()
	} else {
		filePath = args[0]
	}

	if stat, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, FileMode)
		if err != nil {
			log.Fatalf("failed to create directory %s: %s\n", filePath, err)
		}
	} else if !stat.IsDir() {
		log.Fatalf("%s is not a directory\n", filePath)
	}

	if entries, err := os.ReadDir(filePath); err != nil {
		log.Fatalf("failed to read directory %s: %s\n", filePath, err)
	} else if len(entries) > 0 {
		log.Fatalf("Directory %s is not empty\n", filePath)
	}

	yagcDir := path.Join(filePath, ".yagc")
	if err := os.MkdirAll(yagcDir, FileMode); err != nil {
		log.Fatalf("Failed to create directory %s: %s\n", yagcDir, err)
	}

	if head, err := os.Create(path.Join(yagcDir, "HEAD")); err != nil {
		log.Fatalf("Failed to create file %s: %s\n", path.Join(yagcDir, "HEAD"), err)
	} else {
		_, _ = head.WriteString("ref: refs/heads/master")
	}

	repoConfig := config.RepoConfig{Core: config.RepoCore{
		Bare: isBare,
	}}

	if conf, err := yaml.Marshal(repoConfig); err != nil {
		log.Fatalf("Failed to save repo config: %s\n", err)
	} else if err := os.WriteFile(path.Join(yagcDir, "config"), conf, FileMode); err != nil {
		log.Fatalf("Failed to save repo config: %s\n", err)
	}

	for _, file := range []string{"hooks", "into", "objects", "refs", "refs/heads", "refs/tags"} {
		if err := os.MkdirAll(path.Join(yagcDir, file), FileMode); err != nil {
			log.Fatalf("Failed to create directory %s: %s\n", path.Join(yagcDir, file), err)
		}
	}

	log.Printf("Initialized empty repository in %s\n", filePath)
}

func getInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new yagc project",
		Long:  "This command will initialize a new yagc project.",
		Run:   handleInit,
	}

	flags := initCmd.Flags()
	flags.Bool("bare", false, "Create a bare yagc project")
	return initCmd
}
