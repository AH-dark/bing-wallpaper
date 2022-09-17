package util

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func AbsolutePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	base, _ := os.Executable()
	return filepath.Join(filepath.Dir(base), path)
}

// FillSlash 给路径补全`/`
func FillSlash(path string) string {
	if path == "/" {
		return path
	}
	return path + "/"
}

// RemoveSlash 移除路径最后的`/`
func RemoveSlash(path string) string {
	if len(path) > 1 {
		return strings.TrimSuffix(path, "/")
	}
	return path
}

// FormSlash 将path中的反斜杠'\'替换为'/'
func FormSlash(old string) string {
	return path.Clean(strings.ReplaceAll(old, "\\", "/"))
}
