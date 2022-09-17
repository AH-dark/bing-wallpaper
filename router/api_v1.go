package router

import (
	"github.com/AH-dark/bing-wallpaper/middleware"
	"github.com/gin-gonic/gin"
)

func BindApiV1(r *gin.RouterGroup) {
	r.Use(middleware.CORS())
}
