package router

import (
	"compress/gzip"
	"github.com/AH-dark/bing-wallpaper/controller"
	"github.com/AH-dark/bing-wallpaper/middleware"
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Session())

	if conf.StorageConfig.Type == "local" {
		r.GET("/images/*path", middleware.ImagesHandler())
	}

	api := r.Group("/api")
	api.Use(middleware.CORS())
	r.Use(middleware.Gzip(gzip.BestSpeed))

	api.GET("/ping", controller.PingHandler)

	BindApiV1(api.Group("v1"))
	BindApiV2(api.Group("v2"))

	return r
}
