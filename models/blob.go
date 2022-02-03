package models

import (
	"log"
	"os"
	"yagc/util"
)

type BlobObject struct {
	File string `yaml:"file"`
}

func (blob *BlobObject) GetType() ObjectType {
	return Blob
}

func (blob *BlobObject) GetContent() []byte {
	content, err := os.ReadFile(blob.File)
	if err != nil {
		log.Fatalf("Failed to read file %s: %s\n", blob.File, err)
	}
	return util.EncodeObject(blob.GetType(), content)
}

func (blob *BlobObject) GetSha1() (string, []byte) {
	content := blob.GetContent()
	sha1 := util.GetSha1(content)

	return sha1, content
}

func (blob *BlobObject) String() string {
	content, err := os.ReadFile(blob.File)
	if err != nil {
		log.Fatalf("Failed to read file %s: %s\n", blob.File, err)
	}
	return string(content)
}
