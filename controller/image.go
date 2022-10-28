package controller

import (
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func SingleDayImageHandler(c *gin.Context) {
	date, err := time.Parse("20060102", c.Query("date"))
	if err != nil {
		c.Status(400)
		return
	}

	date = util.DayTime(date)

	image, err := model.GetImageByDate(date)
	if err != nil {
		c.Status(500)
		return
	}

	_ = image.ViewUp()

	c.Redirect(302, image.HDUrl)
}

func RandomImageHandler(c *gin.Context) {
	uhd := c.Query("size") == "uhd"

	image, err := model.GetRandomImage()
	if err != nil {
		c.Status(500)
		return
	}

	_ = image.ViewUp()

	if uhd {
		c.Redirect(302, image.UHDUrl)
	} else {
		c.Redirect(302, image.HDUrl)
	}
}
