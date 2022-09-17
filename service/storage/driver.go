package storage

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"io"
	"net/url"
)

type FileInfo struct {
	Name string
	Size int64
	MD5  string
	Key  string
	Url  string
}

type Driver interface {
	Upload(name string, file io.Reader) (*url.URL, error)
}

func NewDriver() (Driver, error) {
	switch conf.StorageConfig.Type {
	case "local":
		return NewLocal(), nil
	case "s3":
		return NewS3(), nil
	}

	return nil, nil
}
