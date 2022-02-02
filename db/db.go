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
	if file, err := os.Create(objectFilePath); err != nil {
		log.Fatalf("Failed to create file %s: %s\n", objectFilePath, err)
	} else {
		file.Write(util.Compress(content))
	}
}
