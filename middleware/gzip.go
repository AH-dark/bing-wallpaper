package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Gzip(level int) gin.HandlerFunc {
	return gzip.Gzip(level)
}
