package model

import (
	"github.com/AH-dark/bing-wallpaper/pkg/cache"
	"strconv"
)

type Setting struct {
	Name string      `gorm:"not null;uniqueIndex" json:"name" xml:"name" yaml:"name"`
	Type SettingType `gorm:"not null;index" json:"type" xml:"type" yaml:"type"`
	Val  string      `gorm:"size:512" json:"val" xml:"val" yaml:"val"`
}

type SettingType string

const (
	SettingTypeSystem  SettingType = "system"
	SettingTypeBasic   SettingType = "basic"
	SettingTypeStorage SettingType = "storage"
	SettingTypeNotify  SettingType = "notify"
)

func GetSetting(name string) (*Setting, error) {
	if v, ok := cache.Get("setting_" + name); ok {
		return v.(*Setting), nil
	}

	var setting Setting
	err := DB.Where("name = ?", name).First(&setting).Error
	if err != nil {
		return nil, err
	}

	_ = cache.Set("setting_"+name, &setting, 0)

	return &setting, err
}

func GetSettingVal(name string) (string, error) {
	d, err := GetSetting(name)
	if err != nil {
		return "", err
	}

	return d.Val, nil
}

func GetSettingValInt(name string) (int, error) {
	d, err := GetSetting(name)
	if err != nil {
		return 0, err
	}

	n, err := strconv.Atoi(d.Val)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func GetSettingValBool(name string) (bool, error) {
	d, err := GetSetting(name)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(d.Val)
}

func GetSettings(key []string) (map[string]string, error) {
	var settings []Setting
	err := DB.Where("name IN ?", key).Find(&settings).Error
	if err != nil {
		return nil, err
	}

	m := make(map[string]string, len(settings))
	for _, v := range settings {
		m[v.Name] = v.Val
	}

	return m, nil
}

func SetSetting(name string, val string) error {
	err := DB.Model(&Setting{}).Where("name = ?", name).Update("val", val).Error
	if err != nil {
		return err
	}

	_ = cache.Delete("setting_" + name)

	return nil
}

func SetSettingValInt(name string, val int) error {
	return SetSetting(name, strconv.Itoa(val))
}

func SetSettingValBool(name string, val bool) error {
	return SetSetting(name, strconv.FormatBool(val))
}

func ListSettingByType(t SettingType) ([]Setting, error) {
	var settings []Setting
	err := DB.Where("type = ?", t).Find(&settings).Error
	return settings, err
}

func ListSetting() ([]Setting, error) {
	var settings []Setting
	err := DB.Find(&settings).Error
	return settings, err
}

func ExistSetting(name string) bool {
	var setting Setting
	err := DB.Where("name = ?", name).First(&setting).Error
	return err == nil
}
