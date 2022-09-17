package bootstrap

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(path string) {
	conf.Init(path)

	if conf.SystemConfig.Debug {
		logger.Level = logrus.DebugLevel
		logger.GlobalLogger = nil
		logger.Log().Debug("Debug mode is on")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
