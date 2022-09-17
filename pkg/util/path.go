package util

import (
	"os"
	"path/filepath"
)

func AbsolutePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	base, _ := os.Executable()
	return filepath.Join(filepath.Dir(base), path)
}
