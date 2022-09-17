package jobs

import (
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/AH-dark/bing-wallpaper/pkg/bing"
	"github.com/AH-dark/bing-wallpaper/service/storage"
	"github.com/AH-dark/logger"
	"net/http"
	"net/url"
	"strings"
)

func FetchBingWallpaper() {
	client := bing.NewClient()

	data, _, err := client.Fetch(0, 7)
	if err != nil {
		logger.Log().Errorf("Fetch bing wallpaper failed: %s", err)
		return
	}

	driver, err := storage.NewDriver()
	if err != nil {
		logger.Log().Errorf("New storage driver failed: %s", err)
		return
	}

	for _, image := range data.Images {
		if model.CheckImageExist(image.StartDate.Time()) {
			continue
		}

		logger.Log().Infof("Fetched bing wallpaper on %s: %s", image.StartDate.Time().Format("2006-01-02"), image.URL)

		q, err := url.Parse("https://www.bing.com/")
		if err != nil {
			logger.Log().Errorf("Parse bing wallpaper url failed: %s", err)
			continue
		}

		q, err = q.Parse(image.URL)
		if err != nil {
			logger.Log().Errorf("Parse bing wallpaper url failed: %s", err)
			continue
		}

		u, _ := url.Parse("https://www.bing.com/")
		u, err = u.Parse(image.URL)
		if err != nil {
			logger.Log().Errorf("Parse bing wallpaper url failed: %s", err)
			continue
		}

		hdRaw, err := http.Get(u.String())
		if err != nil {
			logger.Log().Errorf("Fetch bing wallpaper failed: %s", err)
			continue
		}

		uhdRaw, err := http.Get(strings.Replace(u.String(), "_1920x1080", "_UHD", 1))
		if err != nil {
			logger.Log().Errorf("Fetch bing wallpaper failed: %s", err)
			continue
		}

		hd, err := driver.Upload(q.Query().Get("id"), hdRaw.Body)
		if err != nil {
			logger.Log().Errorf("Upload bing wallpaper failed: %s", err)
			continue
		}

		uhd, err := driver.Upload(strings.Replace(q.Query().Get("id"), "_1920x1080", "_UHD", 1), uhdRaw.Body)
		if err != nil {
			logger.Log().Errorf("Upload bing wallpaper failed: %s", err)
			continue
		}

		img, err := model.CreateImage(model.Image{
			Name:       image.Title,
			StartDate:  image.StartDate.Time(),
			OriginData: image,
			HDUrl:      hd.String(),
			UHDUrl:     uhd.String(),
		})
		if err != nil {
			logger.Log().Errorf("Create image failed: %s", err)
			continue
		}

		logger.Log().Infof("Created image %d: %s (%s)", img.ID, img.Name, img.HDUrl)
	}
}
