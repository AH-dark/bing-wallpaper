package util

import "strings"

func ReplaceString(str string, tasks map[string]string) string {
	for k, v := range tasks {
		str = strings.ReplaceAll(str, k, v)
	}

	return str
}
