package storage

import "github.com/AH-dark/logger"

func TestDriver() {
	d, err := NewDriver()
	if err != nil {
		logger.Log().Fatalf("New storage driver failed: %s", err)
		return
	}

	if err := d.Test(); err != nil {
		logger.Log().Fatalf("Test storage driver failed: %s", err)
		return
	}

	logger.Log().Info("Test storage driver success")
}
