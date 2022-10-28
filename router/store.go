package router

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
	"time"
)

var store persist.CacheStore

func InitStore() {
	if conf.RedisConfig.Server != "" {
		store = persist.NewRedisStore(redis.NewClient(&redis.Options{
			Network:  conf.RedisConfig.Network,
			Addr:     conf.RedisConfig.Server,
			Password: conf.RedisConfig.Password,
			DB:       conf.RedisConfig.DB,
		}))
	} else {
		store = persist.NewMemoryStore(time.Hour)
	}
}

func GetStore() persist.CacheStore {
	return store
}
