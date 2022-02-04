package models

type FileChange int

const (
	FileAdded FileChange = iota
	FileModified
	FileDeleted
)

type Change struct {
	Path         string
	Sha1         string
	OriginalSha1 string
	Kind         FileChange
}
