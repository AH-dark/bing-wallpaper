package cache

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/logger"
)

type Driver interface {
	Get(key string) (interface{}, bool)
	Gets(keys []string, prefix string) (map[string]interface{}, []string)
	Set(key string, val interface{}, ttl int64) error
	Sets(items map[string]interface{}, prefix string) error
	Delete(key string) error
	Deletes(keys []string, prefix string) error
}

var global Driver

func Init() {
	logger.Log().Info("Initialing cache driver...")

	if conf.RedisConfig.Server != "" {
		global = NewRedisDriver(
			10,
			conf.RedisConfig.Network,
			conf.RedisConfig.Server,
			conf.RedisConfig.Password,
			conf.RedisConfig.DB,
		)
	} else {
		global = NewMemoryDriver()
	}
}

func Get(key string) (interface{}, bool) {
	return global.Get(key)
}

func Gets(keys []string, prefix string) (map[string]interface{}, []string) {
	return global.Gets(keys, prefix)
}

func Set(key string, val interface{}, ttl int64) error {
	return global.Set(key, val, ttl)
}

func Sets(items map[string]interface{}, prefix string) error {
	return global.Sets(items, prefix)
}

func Delete(key string) error {
	return global.Delete(key)
}

func Deletes(keys []string, prefix string) error {
	return global.Deletes(keys, prefix)
}
