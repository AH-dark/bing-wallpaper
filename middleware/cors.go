package middleware

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     conf.CORSConfig.AllowOrigins,
		AllowMethods:     conf.CORSConfig.AllowMethods,
		AllowCredentials: conf.CORSConfig.AllowCredentials,
		MaxAge:           time.Duration(conf.CORSConfig.MaxAge) * time.Second,
		AllowHeaders:     conf.CORSConfig.AllowHeaders,
		ExposeHeaders:    conf.CORSConfig.ExposeHeaders,
	})
}
