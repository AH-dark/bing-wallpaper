package model

import "github.com/AH-dark/bing-wallpaper/pkg/conf"

var defaultSettings = []Setting{
	{
		Name: "version",
		Type: SettingTypeSystem,
		Val:  conf.BackendVersion,
	},
	{
		Name: "site_name",
		Type: SettingTypeBasic,
		Val:  "Bing Wallpaper",
	},
	{
		Name: "site_url",
		Type: SettingTypeBasic,
		Val:  "http://localhost:8080",
	},
	{
		Name: "storage_type",
		Type: SettingTypeStorage,
		Val:  "local",
	},
	{
		Name: "storage_local_path",
		Type: SettingTypeStorage,
		Val:  "/data",
	},
	{
		Name: "storage_s3_endpoint",
		Type: SettingTypeStorage,
		Val:  "s3.amazonaws.com",
	},
	{
		Name: "storage_s3_bucket",
		Type: SettingTypeStorage,
		Val:  "bucket",
	},
	{
		Name: "storage_s3_region",
		Type: SettingTypeStorage,
		Val:  "us-east-1",
	},
	{
		Name: "storage_s3_access_key",
		Type: SettingTypeStorage,
		Val:  "",
	},
	{
		Name: "storage_s3_secret_key",
		Type: SettingTypeStorage,
		Val:  "",
	},
	{
		Name: "storage_s3_use_ssl",
		Type: SettingTypeStorage,
		Val:  "true",
	},
	{
		Name: "storage_s3_base_url",
		Type: SettingTypeStorage,
		Val:  "",
	},
	{
		Name: "storage_s3_base_path",
		Type: SettingTypeStorage,
		Val:  "",
	},
}
