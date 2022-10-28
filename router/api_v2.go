package router

import (
	"github.com/AH-dark/bing-wallpaper/controller"
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
	"time"
)

func BindApiV2(r *gin.RouterGroup) {
	r.GET("/single", cache.CacheByRequestURI(store, time.Hour*24), controller.SingleDayImageHandler)
	r.GET("/random", controller.RandomImageHandler)

	data := r.Group("/data")
	{
		data.GET("/list", cache.CacheByRequestURI(store, time.Hour*24), controller.ImageDataListHandler)
		data.GET("/random", controller.RandomImageDataHandler)
	}

	proxy := r.Group("/proxy")
	{
		proxy.GET("/today", cache.CacheByRequestPath(store, time.Hour*24), controller.TodayImageProxyHandler)
		proxy.GET("/single", cache.CacheByRequestURI(store, time.Hour*24), controller.SingleImageProxyHandler)
	}
}
