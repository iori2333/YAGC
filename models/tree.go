package models

import (
	"bytes"
	"log"
	"yagc/util"
)

// TreeObject is a tree object that contains a list of entries, which
// are either other trees or blobs
type TreeObject struct {
	Id      string  `yaml:"id"`      // Sha1 of tree
	Entries []Entry `yaml:"entries"` // List of entries
}

func (object *TreeObject) GetType() ObjectType {
	return Tree
}

func (object *TreeObject) GetContent() []byte {
	buf := bytes.Buffer{}
	for _, entry := range object.Entries {
		buf.WriteString(entry.String())
	}
	return util.EncodeObject(object.GetType(), buf.Bytes())
}

func (object *TreeObject) GetSha1() (string, []byte) {
	content := object.GetContent()
	if object.Id == "" {
		sha1 := util.GetSha1(content)
		object.Id = sha1
	}
	return object.Id, content
}

func (object *TreeObject) Parse(content []byte) *TreeObject {
	entries := make([]Entry, 0)
	for _, line := range bytes.Split(content, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		entry := bytes.SplitN(line, []byte(" "), 4)
		if len(entry) != 4 {
			log.Fatalf("Failed to parse tree entry: %s\n", entry)
		}
		entries = append(entries, Entry{
			Id:   string(entry[0]),
			Type: string(entry[1]),
			Mode: string(entry[2]),
			File: string(entry[3]),
		})
	}
	object.Entries = entries
	return object
}

func (object *TreeObject) Exists(sha1 string) bool {
	for _, entry := range object.Entries {
		if entry.Id == sha1 {
			return true
		}
	}
	return false
}

func (object *TreeObject) AddEntry(entry Entry) {
	object.Entries = append(object.Entries, entry)
}
