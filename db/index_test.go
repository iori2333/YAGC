package db

import (
	"os"
	"testing"
)

func TestIndexing(t *testing.T) {
	entries, _ := os.ReadDir(".")
	names := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name()
	}
	CreateIndex(names)
}
