package controller

import (
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ImageDataListHandler(c *gin.Context) {
	start, err := time.Parse("20060102", c.Query("start"))
	if err != nil {
		c.Status(400)
		return
	}

	end, err := time.Parse("20060102", c.Query("end"))
	if err != nil {
		c.Status(400)
		return
	}

	images, err := model.GetImageList(start, end)
	if err != nil {
		c.Status(500)
		return
	}

	if len(images) == 0 {
		c.Status(404)
		return
	}

	l := end.Sub(start).Hours()/24 + 1
	if len(images) != int(l) {
		c.Header("X-Missing", strconv.Itoa(int(l)-len(images)))
	}

	go func() {
		for _, image := range images {
			_ = image.ViewUp()
		}
	}()

	c.JSON(http.StatusOK, images)
}

func RandomImageDataHandler(c *gin.Context) {
	image, err := model.GetRandomImage()
	if err != nil {
		c.Status(500)
		return
	}

	_ = image.ViewUp()

	c.JSON(http.StatusOK, image)
}
