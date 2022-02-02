package db

import (
	"log"
	"os"
	"path"
	"yagc/util"
)

func Write(id string, content []byte) {
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

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file %s: %s\n", objectFilePath, err)
		}
	}()

	if err != nil {
		log.Fatalf("Failed to create file %s: %s\n", objectFilePath, err)
	} else {
		file.Write(util.Compress(content))
	}
}

func Index() {
	log.Fatalln("Not implemented")
}

func Find(id string) []byte {
	prefix, suffix := id[:2], id[2:]
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	objectPath := path.Join(root, ".yagc", "objects", prefix, suffix)
	content, err := os.ReadFile(objectPath)
	if err != nil {
		log.Fatalln("Failed to read object", id)
	}
	content = util.Decompress(content)

	return content
}
