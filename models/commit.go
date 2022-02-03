package models

import (
	"gopkg.in/yaml.v2"
	"log"
	"yagc/util"
)

// CommitObject is a representation of a commit object. It contains the
// basic information about a commit and its TreeObject.
type CommitObject struct {
	Id         string `yaml:"id"`          // Sha1 of the commit
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
	if commit.Id == "" {
		sha1 := util.GetSha1(content)
		commit.Id = sha1
	}
	return commit.Id, content
}

func (commit *CommitObject) Parse(content []byte) *CommitObject {
	err := yaml.Unmarshal(content, commit)
	if err != nil {
		log.Fatalf("Unable to parse commit object: %v", err)
	}
	return commit
}
