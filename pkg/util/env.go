package util

import (
	"os"
	"strconv"
	"strings"
)

func EnvStr(name, defaultValue string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}

	return defaultValue
}

func EnvInt(name string, defaultValue int) int {
	if value := os.Getenv(name); value != "" {
		n, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}

		return n
	}

	return defaultValue
}

func EnvArr(name string, defaultValue []string) []string {
	if value := os.Getenv(name); value != "" {
		return strings.Split(value, ",")
	}

	return defaultValue
}
