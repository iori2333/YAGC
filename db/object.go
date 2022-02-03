package db

import (
	"log"
	"os"
	"path"
	"strings"
	"yagc/util"
)

func WriteObject(id string, content []byte) {
	if ExistsObject(id) {
		return
	}

	prefix, suffix := id[:2], id[2:]
	root, ok := util.GetRepoRoot()

	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	objectPath := path.Join(root, ".yagc", "objects", prefix)
	if err := os.MkdirAll(objectPath, 0644); err != nil {
		log.Fatalf("Failed to create directory %s: %s\n", objectPath, err)
	}

	objectFilePath := path.Join(objectPath, suffix)
	file, err := os.Create(objectFilePath)

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file %s: %s\n", objectFilePath, err)
		}
	}(file)

	if err != nil {
		log.Fatalf("Failed to create file %s: %s\n", objectFilePath, err)
	}

	if _, err := file.Write(util.Compress(content)); err != nil {
		log.Fatalf("Failed to write to file %s: %s\n", objectFilePath, err)
	}
}

func FindObject(id string) []byte {
	if len(id) <= 2 {
		log.Fatalf("Invalid object name %s\n", id)
	}

	prefix, suffix := id[:2], id[2:]
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	objectPath := path.Join(root, ".yagc", "objects", prefix)
	entries, _ := os.ReadDir(objectPath)

	var objectFilePath string
	for _, file := range entries {
		if strings.HasPrefix(file.Name(), suffix) {
			if objectFilePath != "" {
				log.Fatalf("Multiple objects with name %s\n", id)
			}
			objectFilePath = path.Join(objectPath, file.Name())
		}
	}

	if objectFilePath == "" {
		log.Fatalf("Failed to find object %s\n", id)
	}

	content, err := os.ReadFile(objectFilePath)
	if err != nil {
		log.Fatalln("Failed to read object", id)
	}

	return util.Decompress(content)
}

func ExistsObject(sha1 string) bool {
	prefix, suffix := sha1[:2], sha1[2:]
	root, ok := util.GetRepoRoot()

	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	objectPath := path.Join(root, ".yagc", "objects", prefix, suffix)
	if _, err := os.Stat(objectPath); os.IsNotExist(err) {
		return false
	}
	return true
}
