package middleware

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"strconv"
)

var store sessions.Store

func Session() gin.HandlerFunc {
	var err error
	if conf.RedisConfig.Server != "" {
		store, err = redis.NewStoreWithDB(
			10,
			"tcp",
			conf.RedisConfig.Server,
			conf.RedisConfig.Password,
			strconv.Itoa(conf.RedisConfig.DB),
			[]byte(conf.SystemConfig.SessionSecret),
		)
		if err != nil {
			logger.Log().Panicf("Redis session store error: %s", err)
		}
	} else {
		store = cookie.NewStore([]byte(conf.SystemConfig.SessionSecret))
	}

	return sessions.Sessions("go-session", store)
}
