package main

import (
	"flag"
	"github.com/AH-dark/bing-wallpaper/bootstrap"
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/router"
	"github.com/AH-dark/logger"
)

var (
	configPath = "config.ini"
)

func init() {
	flag.StringVar(&configPath, "c", configPath, "config file path")
	flag.Parse()

	bootstrap.InitApplication(configPath)
}

func main() {
	r := router.Init()

	logger.Log().Infof("Starting server at %s", conf.SystemConfig.Listen)
	err := r.Run(conf.SystemConfig.Listen)
	if err != nil {
		logger.Log().Panicf("Run server error: %v", err)
		return
	}
}
