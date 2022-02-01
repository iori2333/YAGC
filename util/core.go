package util

import (
	"os"
	"path"
	"runtime"
)

func GetRepoRoot() (string, bool) {
	_, caller, _, _ := runtime.Caller(1)
	workDir := path.Dir(caller)

	for workDir != "/" && workDir != "." {
		if f, err := os.Stat(path.Join(workDir, ".yagc")); err == nil && f.IsDir() {
			return workDir, true
		} else {
			workDir = path.Dir(workDir)
		}
	}
	return "", false
}
