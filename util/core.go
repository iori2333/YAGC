package util

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func GetRepoRoot() (string, bool) {
	workDir, _ := os.Getwd()
	workDir, _ = filepath.Abs(workDir)

	for workDir != "/" && workDir != "." {
		if f, err := os.Stat(path.Join(workDir, ".yagc")); err == nil && f.IsDir() {
			return workDir, true
		} else {
			workDir = path.Dir(workDir)
		}
	}
	return "", false
}

func Compress(content []byte) []byte {
	buf := bytes.Buffer{}
	writer := zlib.NewWriter(&buf)

	if _, err := writer.Write(content); err != nil {
		log.Fatalf("Failed to write to zlib writer: %s\n", err)
	}

	writer.Close()
	return buf.Bytes()
}

func Decompress(content []byte) []byte {
	reader, err := zlib.NewReader(bytes.NewReader(content))
	if err != nil {
		log.Fatalf("Failed to create zlib reader: %s\n", err)
	}
	defer reader.Close()

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(reader); err != nil {
		log.Fatalf("Failed to read from zlib reader: %s\n", err)
	}

	return buf.Bytes()
}

func GetSha1(content []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(content))
}
