package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
	"yagc/util"
)

// CommitObject is a representation of a commit object. It contains the
// basic information about a commit and its TreeObject.
type CommitObject struct {
	Message    string `yaml:"message"`     // Commit message
	Author     string `yaml:"author"`      // Author of the commit
	Committer  string `yaml:"committer"`   // Committer of the commit
	Date       string `yaml:"date"`        // Date of the commit
	Parent     string `yaml:"parent"`      // The commit this commit is based on
	TreeObject string `yaml:"tree_object"` // Sha1 of the related tree object
}

func (commit *CommitObject) GetType() ObjectType {
	return Commit
}

func (commit *CommitObject) GetContent() []byte {
	out, err := yaml.Marshal(commit)
	if err != nil {
		log.Fatalf("Unable to create commit object: %v", err)
	}

	return util.EncodeObject(commit.GetType(), out)
}

func (commit *CommitObject) GetSha1() (string, []byte) {
	content := commit.GetContent()
	sha1 := util.GetSha1(content)

	return sha1, content
}

func (commit *CommitObject) Parse(content []byte) *CommitObject {
	objType, _, content, err := util.DecodeObject(content)
	if err != nil || objType != Commit {
		log.Fatalf("Unable to parse commit object: %v", err)
	}

	if err := yaml.Unmarshal(content, commit); err != nil {
		log.Fatalf("Unable to parse commit object: %v", err)
	}

	return commit
}

func (commit *CommitObject) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("tree %s\n", commit.TreeObject))
	if commit.Parent != "" {
		builder.WriteString(fmt.Sprintf("parent %s\n", commit.Parent))
	}
	builder.WriteString(fmt.Sprintf("author %s\n", commit.Author))
	builder.WriteString(fmt.Sprintf("committer %s\n", commit.Committer))
	builder.WriteString(fmt.Sprintf("date %s\n", commit.Date))
	builder.WriteString(fmt.Sprintf("\n%s\n", commit.Message))

	return builder.String()
}
