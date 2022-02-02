package models

type ObjectType = string

const (
	Blob   ObjectType = "blob"
	Commit ObjectType = "commit"
	Tree   ObjectType = "tree"
)

type Object interface {
	GetType() ObjectType
	GetContent() []byte
	GetSha1() (string, []byte)
}
