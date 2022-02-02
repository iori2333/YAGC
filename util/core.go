package util

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

var workDir string

func init() {
	workDir, _ = os.Getwd()
	workDir, _ = filepath.Abs(workDir)

	for workDir != "/" && workDir != "." {
		f, err := os.Stat(path.Join(workDir, ".yagc"))
		if err == nil && f.IsDir() {
			return
		}
		workDir = path.Dir(workDir)
	}

	if workDir == "/" || workDir == "." {
		workDir = ""
	}
}

func GetRepoRoot() (string, bool) {
	return workDir, workDir != ""
}

func Compress(content []byte) []byte {
	buf := bytes.Buffer{}
	writer := zlib.NewWriter(&buf)

	if _, err := writer.Write(content); err != nil {
		log.Fatalf("Failed to write to zlib writer: %s\n", err)
	}

	err := writer.Close()
	if err != nil {
		log.Fatalf("Failed to close zlib writer: %s\n", err)
	}
	return buf.Bytes()
}

func Decompress(content []byte) []byte {
	reader, err := zlib.NewReader(bytes.NewReader(content))
	if err != nil {
		log.Fatalf("Failed to create zlib reader: %s\n", err)
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			log.Fatalf("Failed to close zlib reader: %s\n", err)
		}
	}(reader)

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(reader); err != nil {
		log.Fatalf("Failed to read from zlib reader: %s\n", err)
	}

	return buf.Bytes()
}

func GetSha1(content []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(content))
}
