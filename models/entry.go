package models

import (
	"fmt"
)

// Entry is a node in the tree that points to any type of object
//
// If Entry has type Tree, it might point to other entries;
//
// If Entry has type Commit, it stores information about a commit,
// which contains a corresponding TreeObject;
//
// Else it points to a Blob that stores a regular file
type Entry struct {
	Id   string `yaml:"id"`   // Sha1 of the object
	Type string `yaml:"type"` // Type of the object
	Mode string `yaml:"mode"` // File mode of pointed object
	File string `yaml:"file"` // File name of pointed object
}

func (entry *Entry) String() string {
	return fmt.Sprintf("%s %s %s %s", entry.Id, entry.Type, entry.Mode, entry.File)
}
