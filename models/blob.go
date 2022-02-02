package models

import (
	"fmt"
	"log"
	"os"
	"yagc/util"
)

type BlobObject struct {
	Id   string `yaml:"id"`
	File string `yaml:"file"`
}

func (blob *BlobObject) GetType() FileType {
	return Blob
}

func (blob *BlobObject) GetContent() []byte {
	content, err := os.ReadFile(blob.File)
	if err != nil {
		log.Fatalf("Failed to read file %s: %s\n", blob.File, err)
	}
	header := fmt.Sprintf("%s %d\u0000", blob.GetType(), len(content))
	content = append([]byte(header), content...)
	return content
}

func (blob *BlobObject) GetSha1() (string, []byte) {
	content := blob.GetContent()
	if blob.Id == "" {
		sha1 := util.GetSha1(content)
		blob.Id = sha1
	}
	return blob.Id, content
}
