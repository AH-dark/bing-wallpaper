package router

import (
	"github.com/AH-dark/bing-wallpaper/controller"
	"github.com/gin-gonic/gin"
)

func BindApiV2(r *gin.RouterGroup) {
	r.GET("/single", controller.SingleDayImageHandler)
	r.GET("/random", controller.RandomImageHandler)

	data := r.Group("/data")
	{
		data.GET("/list", controller.ImageDataListHandler)
		data.GET("/random", controller.RandomImageDataHandler)
	}
}
