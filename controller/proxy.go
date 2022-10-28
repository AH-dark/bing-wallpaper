package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"github.com/AH-dark/logger"
	"github.com/gin-gonic/gin"
)

func TodayImageProxyHandler(c *gin.Context) {
	date := util.TodayBegin()

	image, err := model.GetImageByDate(date)
	if err != nil {
		logger.Log().Errorf("GetImageByDate error: %s", err)
		c.String(500, "Internal Server Error")
		return
	}

	_ = image.ViewUp()

	resp, err := http.Get(image.HDUrl)
	if err != nil || resp.StatusCode != 200 {
		logger.Log().Errorf("Get image error: %s", err)
		c.String(500, "Internal Server Error")
		return
	}

	c.DataFromReader(200, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, map[string]string{
		"Cache-Control":       "public, max-age=86400, immutable",
		"Content-Disposition": fmt.Sprintf("inline; filename=\"%s\"", image.Name+".jpg"),
		"Content-Length":      strconv.Itoa(int(resp.ContentLength)),
		"Expires":             date.AddDate(0, 0, 1).Format(time.RFC1123),
		"Last-Modified":       date.Format(time.RFC1123),
		"Link":                fmt.Sprintf("<%s>; rel=\"canonical\"", image.HDUrl),
		"Vary":                "Accept-Encoding",

		"X-Proxy":       "BingWallpaper/" + conf.BackendVersion,
		"X-Image-Date":  date.Format(time.RFC3339),
		"X-Image-Url":   image.HDUrl,
		"X-Image-View":  fmt.Sprintf("%d", image.View),
		"X-Image-Title": image.Name,
	})
}

func SingleImageProxyHandler(c *gin.Context) {
	date, err := time.Parse("20060102", c.Query("date"))
	if err != nil {
		c.String(400, "Bad Request, date format error")
		return
	}

	date = util.DayTime(date)

	image, err := model.GetImageByDate(date)
	if err != nil {
		logger.Log().Errorf("GetImageByDate error: %s", err)
		c.String(500, "Internal Server Error")
		return
	}

	_ = image.ViewUp()

	resp, err := http.Get(image.HDUrl)
	if err != nil || resp.StatusCode != 200 {
		logger.Log().Errorf("Get image error: %s", err)
		c.String(500, "Internal Server Error")
		return
	}

	c.DataFromReader(200, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, map[string]string{
		"Cache-Control":       "public, max-age=86400, immutable",
		"Content-Disposition": fmt.Sprintf("inline; filename=\"%s\"", image.Name+".jpg"),
		"Content-Length":      strconv.Itoa(int(resp.ContentLength)),
		"Expires":             date.AddDate(0, 0, 1).Format(time.RFC1123),
		"Last-Modified":       date.Format(time.RFC1123),
		"Link":                fmt.Sprintf("<%s>; rel=\"canonical\"", image.HDUrl),
		"Vary":                "Accept-Encoding",

		"X-Proxy":       "BingWallpaper/" + conf.BackendVersion,
		"X-Image-Date":  date.Format(time.RFC3339),
		"X-Image-Url":   image.HDUrl,
		"X-Image-View":  fmt.Sprintf("%d", image.View),
		"X-Image-Title": image.Name,
	})
}
