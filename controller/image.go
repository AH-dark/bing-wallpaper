package controller

import (
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/gin-gonic/gin"
	"time"
)

func SingleDayImageHandler(c *gin.Context) {
	date, err := time.Parse("20060102", c.Query("date"))
	if err != nil {
		c.Status(400)
		return
	}

	image, err := model.GetImageByDate(date)
	if err != nil {
		c.Status(500)
		return
	}

	c.Redirect(302, image.HDUrl)
}
