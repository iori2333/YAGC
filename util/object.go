package util

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
)

func Compress(content []byte) []byte {
	buf := bytes.Buffer{}
	writer := zlib.NewWriter(&buf)
	writer.Write(content)
	writer.Close()
	return buf.Bytes()
}

func GetSha1(content []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(content))
}
