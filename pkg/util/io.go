package util

import (
	"os"
	"path/filepath"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func CreateNestFile(path string) (*os.File, error) {
	path = AbsolutePath(path)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}

	return os.Create(path)
}
