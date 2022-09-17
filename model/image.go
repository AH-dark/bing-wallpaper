package model

import (
	"github.com/AH-dark/bing-wallpaper/pkg/cache"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Image struct {
	gorm.Model `json:"-" xml:"-" yaml:"-"`

	StartDate time.Time `gorm:"not null;unique" json:"start_date" xml:"start_date" yaml:"start_date"`

	Name       string      `json:"name" xml:"name" yaml:"name"`
	OriginData interface{} `gorm:"serializer:json;size:512" json:"origin_data" xml:"origin_data" yaml:"origin_data"`
	HDUrl      string      `gorm:"type:varchar(255);not null;unique" json:"hd_url" xml:"hd_url" yaml:"hd_url"`
	UHDUrl     string      `gorm:"type:varchar(255);not null;unique" json:"uhd_url" xml:"uhd_url" yaml:"uhd_url"`

	View int64 `gorm:"type:bigint;not null;default:0" json:"view" xml:"view" yaml:"view"`
	Love int64 `gorm:"type:bigint;not null;default:0" json:"love" xml:"love" yaml:"love"`
}

func CreateImage(image Image) (*Image, error) {
	err := DB.Create(&image).Error
	if err != nil {
		return nil, err
	}

	_ = cache.Delete("count_images")

	return &image, nil
}

func GetImageById(id uint) (*Image, error) {
	key := "image_" + strconv.Itoa(int(id))
	if v, ok := cache.Get(key); ok {
		return v.(*Image), nil
	}

	var image Image
	err := DB.Where("id = ?", id).First(&image).Error
	if err != nil {
		return nil, err
	}

	_ = cache.Set(key, &image, 0)

	return &image, nil
}

func GetImageList(startDate, endDate time.Time) ([]Image, error) {
	var images []Image
	err := DB.Where("start_date >= ? AND start_date <= ?", startDate, endDate).Find(&images).Error
	return images, err
}

func GetLatestImage() (*Image, error) {
	var image Image
	err := DB.Order("start_date DESC").First(&image).Error
	return &image, err
}

func GetLatestImages(num int) ([]Image, error) {
	var images []Image
	err := DB.Order("start_date DESC").Limit(num).Find(&images).Error
	return images, err
}

func GetImageListByDateRange(startDate, endDate time.Time) ([]Image, error) {
	var images []Image
	err := DB.Where("start_date >= ? AND start_date <= ?", startDate, endDate).Find(&images).Error
	return images, err
}

func GetImageByDate(date time.Time) (*Image, error) {
	var image Image
	err := DB.Where("start_date = ?", date).First(&image).Error
	return &image, err
}

func CountImages() int64 {
	if v, ok := cache.Get("count_images"); ok {
		return v.(int64)
	}

	var count int64 = 0
	DB.Model(&Image{}).Count(&count)

	_ = cache.Set("count_images", count, 0)

	return count
}

func GetRandomImage() (*Image, error) {
	num := util.RandInt(0, int(CountImages()))
	var image Image
	err := DB.Order("start_date DESC").Offset(num).First(&image).Error
	return &image, err
}

func CheckImageExist(startDate time.Time) bool {
	var count int64 = 0
	DB.Model(&Image{}).Where("start_date = ?", startDate).Count(&count)
	return count > 0
}

func (image *Image) delete() error {
	return DB.Delete(image).Error
}

func (image *Image) Delete() error {
	return image.delete()
}

func (image *Image) ViewUp() error {
	image.View++
	return image.update()
}

func (image *Image) LoveUp() error {
	image.Love++
	return image.update()
}

func (image *Image) update() error {
	return DB.Save(image).Error
}
