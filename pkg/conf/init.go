package conf

import (
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"github.com/AH-dark/logger"
	"gopkg.in/ini.v1"
)

func Init(path string) {
	path = util.AbsolutePath(path)

	if !util.Exist(path) {
		f, err := util.CreateNestFile(path)
		if err != nil {
			logger.Log().Panicf("Create config file error: %v", err)
			return
		}

		_, err = f.WriteString(util.ReplaceString(defaultConfig, map[string]string{
			"{{SessionSecret}}": util.RandString(32),
		}))
		if err != nil {
			logger.Log().Panicf("Write config file error: %v", err)
			return
		}

		_ = f.Close()
	}

	cfg, err := ini.Load(path)
	if err != nil {
		logger.Log().Panicf("Load config file error: %v", err)
		return
	}

	mapping := map[string]interface{}{
		"System":   SystemConfig,
		"Database": DatabaseConfig,
		"Redis":    RedisConfig,
		"CORS":     CORSConfig,
	}

	for k, v := range mapping {
		err = cfg.Section(k).MapTo(v)
		if err != nil {
			logger.Log().Panicf("Map config file error: %v", err)
			return
		}
	}

	logger.Log().Infof("Load config file success: %s", path)
}
