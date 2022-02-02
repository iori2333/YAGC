package util

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

func EncodeObject(objType string, content []byte) []byte {
	header := fmt.Sprintf("%s %d ", objType, len(content))
	content = append([]byte(header), content...)
	return content
}

func DecodeObject(content []byte) (string, int, []byte, error) {
	header := bytes.SplitN(content, []byte(" "), 3)
	if len(header) != 3 {
		return "", 0, nil, errors.New("failed to parse object header")
	}

	size, err := strconv.ParseInt(string(header[1]), 10, 0)
	if err != nil {
		return "", 0, nil, errors.New("failed to parse object size")
	}

	objType, content := string(header[0]), header[2]
	return objType, int(size), content, nil
}
