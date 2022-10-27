package storage

import (
	"fmt"
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"io"
	"net/url"
	"os"
	"strings"
)

type LocalImpl struct {
}

func (l *LocalImpl) Upload(name string, file io.Reader) (*url.URL, error) {
	siteUrl, err := model.GetSettingVal("site_url")
	if err != nil {
		return nil, err
	}

	f, err := util.CreateNestFile(util.AbsolutePath(fmt.Sprintf("%s/%s", strings.Trim(conf.StorageConfig.BasePath, "/"), name)))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(siteUrl)
	if err != nil {
		return nil, err
	}

	u, _ = u.Parse("/images/")
	u, err = u.Parse(conf.StorageConfig.BasePath)
	if err != nil {
		return nil, err
	}

	u, err = u.Parse(name)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (l *LocalImpl) Test() error {
	if !util.Exist(util.AbsolutePath(conf.StorageConfig.BasePath)) {
		if err := os.MkdirAll(util.AbsolutePath(conf.StorageConfig.BasePath), os.ModePerm); err != nil {
			return err
		}
	}

	f, err := util.CreateNestFile(util.AbsolutePath(fmt.Sprintf("%s/test.txt", strings.Trim(conf.StorageConfig.BasePath, "/"))))
	if err != nil {
		return err
	}

	if _, err = f.WriteString("test"); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Remove(util.AbsolutePath(fmt.Sprintf("%s/test.txt", strings.Trim(conf.StorageConfig.BasePath, "/")))); err != nil {
		return err
	}

	return nil
}

func NewLocal() Driver {
	return &LocalImpl{}
}
