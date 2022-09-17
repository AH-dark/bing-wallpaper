package cron

import (
	"github.com/AH-dark/bing-wallpaper/pkg/cron/jobs"
	"github.com/AH-dark/logger"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func Init() {
	if c == nil {
		c = cron.New()
	}

	id, err := c.AddFunc("@daily", jobs.FetchBingWallpaper)
	if err != nil {
		logger.Log().Warningf("Add cron job %d failed: %s", id, err)
		return
	}

	go jobs.FetchBingWallpaper()

	c.Start()
}

func Reset() {
	c.Stop()
	c = nil
	Init()
}
