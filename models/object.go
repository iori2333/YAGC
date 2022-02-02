package models

type FileType = string

const (
	Blob   FileType = "blob"
	Commit FileType = "commit"
	Tree   FileType = "tree"
)

type Object interface {
	GetType() FileType
	GetContent() []byte
	GetSha1() string
}
