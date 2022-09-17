package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceString(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal("hello world", ReplaceString("hallo world", map[string]string{
		"hallo": "hello",
	}))
}
