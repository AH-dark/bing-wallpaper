package router

import (
	"github.com/AH-dark/bing-wallpaper/controller"
	"github.com/gin-gonic/gin"
)

func BindApiV2(r *gin.RouterGroup) {
	r.GET("/single", controller.SingleDayImageHandler)
}
