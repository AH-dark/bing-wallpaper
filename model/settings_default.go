package model

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
)

var defaultSettings = []Setting{
	{Name: "version", Type: SettingTypeSystem, Val: conf.BackendVersion},
	{Name: "site_name", Type: SettingTypeBasic, Val: "Bing Wallpaper"},
	{Name: "site_url", Type: SettingTypeBasic, Val: "http://localhost:8080"},
}
