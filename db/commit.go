package db

import (
	"bytes"
	"log"
	"os"
	"path"
	"strconv"
	"time"
	"yagc/config"
	"yagc/models"
	"yagc/util"
)

func GetRef() string {
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}
	refByte, err := os.ReadFile(path.Join(root, ".yagc", "HEAD"))
	if err != nil {
		log.Fatalln("Failed to read HEAD")
	}
	refs := bytes.SplitN(refByte, []byte(" "), 2)
	if len(refs) != 2 {
		log.Fatalln("Failed to parse HEAD")
	}
	return string(refs[1])
}

func Commit(message string) {
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	ref := GetRef()
	conf := config.App()
	if !conf.User.Valid() {
		log.Fatalf("User name and email are required to commit")
	}

	index := ReadIndex()

	sha1, content := index.GetSha1()
	WriteObject(sha1, content)

	commit := models.CommitObject{
		Message:    message,
		Author:     conf.User.String(),
		Committer:  conf.User.String(),
		Date:       strconv.FormatInt(time.Now().Unix(), 10),
		Parent:     LastCommit(),
		TreeObject: sha1,
	}

	sha1, content = commit.GetSha1()
	WriteObject(sha1, content)

	if err := os.WriteFile(path.Join(root, ".yagc", ref), []byte(sha1), 0644); err != nil {
		log.Fatalln("Failed to update ref")
	}
}

func LastCommit() string {
	root, ok := util.GetRepoRoot()
	if !ok {
		log.Fatalln("Failed to get repo root")
	}

	ref := GetRef()

	curr, err := os.ReadFile(path.Join(root, ".yagc", ref))
	if err != nil {
		return ""
	}

	return string(curr)
}
