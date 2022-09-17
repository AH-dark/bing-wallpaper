package model

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/logger"
)

func migrate() {
	if !checkMigrate() {
		logger.Log().Info("No need to migrate")
		return
	}

	logger.Log().Info("Migrating...")

	err := DB.AutoMigrate(&Image{}, &Setting{})
	if err != nil {
		logger.Log().Panicf("Migrate error: %v", err)
		return
	}

	addDefaultSettings()

	logger.Log().Info("Migrate succeed")
}

func checkMigrate() bool {
	if conf.SystemConfig.Debug {
		return true
	}

	if v, err := GetSetting("version"); err != nil || v.Val != conf.BackendVersion {
		return true
	}

	return false
}

func addDefaultSettings() {
	logger.Log().Info("Adding default settings...")

	for _, setting := range defaultSettings {
		if ExistSetting(setting.Name) {
			continue
		}

		err := DB.Create(&Setting{
			Name: setting.Name,
			Type: setting.Type,
			Val:  setting.Val,
		}).Error
		if err != nil {
			logger.Log().Errorf("Add default settings error: %v", err)
			continue
		}
	}
}
