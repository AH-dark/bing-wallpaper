package util

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func CreateNestFile(path string) (*os.File, error) {
	path = AbsolutePath(path)

	err := os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return nil, err
	}

	return os.Create(path)
}

func GzipCompress(reader io.Reader) (io.Reader, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	defer zw.Close()

	_, err := io.Copy(zw, reader)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
