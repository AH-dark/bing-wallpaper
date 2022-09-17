package middleware

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"github.com/AH-dark/logger"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ImagesHandler() gin.HandlerFunc {
	ignoreFunc := func(c *gin.Context) {
		c.Next()
	}

	dir := util.AbsolutePath(conf.StorageConfig.BasePath)

	return func(c *gin.Context) {
		logger.Log().Debugf("Request path: %s", c.Request.URL.Path)

		if c.Request.URL.Path == "/images" {
			ignoreFunc(c)
			return
		}

		if c.Request.URL.Path == "/images/" {
			ignoreFunc(c)
			return
		}

		if !util.Exist(filepath.Join(dir, c.Request.URL.Path)) {
			c.Status(http.StatusNotFound)
		}

		f, err := os.Open(filepath.Join(dir, strings.TrimLeft(c.Request.URL.Path, "/images/")))
		if err != nil {
			logger.Log().Errorf("Failed to open file: %s", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			logger.Log().Errorf("Failed to get file stat: %s", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.DataFromReader(http.StatusOK, stat.Size(), mime.TypeByExtension(filepath.Ext(c.Request.URL.Path)), f, map[string]string{})
		c.Abort()
	}
}
