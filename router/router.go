package router

import (
	"compress/gzip"
	"github.com/AH-dark/bing-wallpaper/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Session())

	api := r.Group("/api")
	api.Use(middleware.CORS())
	r.Use(middleware.Gzip(gzip.BestSpeed))

	BindApiV1(api.Group("v1"))

	return r
}
