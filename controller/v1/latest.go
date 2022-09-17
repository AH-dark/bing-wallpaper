package v1

import (
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/gin-gonic/gin"
)

func LatestImageHandler(c *gin.Context) {
	image, err := model.GetLatestImage()
	if err != nil {
		c.Status(500)
		return
	}

	c.Redirect(302, image.HDUrl)
}
