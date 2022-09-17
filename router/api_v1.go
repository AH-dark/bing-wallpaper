package router

import (
	"github.com/AH-dark/bing-wallpaper/controller/v1"
	"github.com/AH-dark/bing-wallpaper/middleware"
	"github.com/gin-gonic/gin"
)

func BindApiV1(r *gin.RouterGroup) {
	r.Use(middleware.CORS())

	r.GET("/new", v1.LatestImageHandler)
}
