package models

import (
	"gopkg.in/yaml.v2"
	"log"
	"strings"
	"yagc/util"
)

// TreeObject is a tree object that contains a list of entries, which
// are either other trees or blobs
type TreeObject struct {
	Entries []Entry `yaml:"entries"` // List of entries
}

func (tree *TreeObject) GetType() ObjectType {
	return Tree
}

func (tree *TreeObject) GetContent() []byte {
	out, err := yaml.Marshal(tree)
	if err != nil {
		log.Fatalf("Error marshalling tree object: %s", err)
	}
	return util.EncodeObject(tree.GetType(), out)
}

func (tree *TreeObject) GetSha1() (string, []byte) {
	content := tree.GetContent()
	sha1 := util.GetSha1(content)

	return sha1, content
}

func (tree *TreeObject) Parse(content []byte) *TreeObject {
	err := yaml.Unmarshal(content, tree)
	if err != nil {
		log.Fatalf("Error unmarshalling tree object: %s", err)
	}
	return tree
}

func (tree *TreeObject) Exists(sha1 string) bool {
	for _, entry := range tree.Entries {
		if entry.Id == sha1 {
			return true
		}
	}
	return false
}

func (tree *TreeObject) AddEntry(entry Entry) {
	tree.Entries = append(tree.Entries, entry)
}

func (tree *TreeObject) String() string {
	builder := strings.Builder{}
	for _, entry := range tree.Entries {
		builder.WriteString(entry.String())
	}
	return builder.String()
}
