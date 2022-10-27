package bootstrap

import (
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/AH-dark/bing-wallpaper/pkg/cache"
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/pkg/cron"
	"github.com/AH-dark/bing-wallpaper/service/storage"
	"github.com/AH-dark/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(path string, skip bool) {
	if !skip {
		conf.Init(path)
	}

	if conf.SystemConfig.Debug {
		logger.Level = logrus.DebugLevel
		logger.GlobalLogger = nil
		logger.Log().Debug("Debug mode is on")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	cache.Init()

	model.Init()

	cron.Init()

	storage.TestDriver()
}
