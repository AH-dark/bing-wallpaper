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

	Uri string `gorm:"type:varchar(255);not null;uniqueIndex" json:"uri" xml:"uri" yaml:"uri"`

	View int64 `gorm:"type:bigint;not null;default:0" json:"view" xml:"view" yaml:"view"`
	Love int64 `gorm:"type:bigint;not null;default:0" json:"love" xml:"love" yaml:"love"`
}

func CreateImage(uri string) (*Image, error) {
	image := &Image{
		Uri: uri,
	}
	err := DB.Create(image).Error
	if err != nil {
		return nil, err
	}

	_ = cache.Delete("count_images")

	return image, nil
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
	err := DB.Where("created_at >= ? AND created_at <= ?", startDate, endDate).Find(&images).Error
	return images, err
}

func GetLatestImage() (*Image, error) {
	var image Image
	err := DB.Order("created_at DESC").First(&image).Error
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
	err := DB.Offset(num).First(&image).Error
	return &image, err
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
