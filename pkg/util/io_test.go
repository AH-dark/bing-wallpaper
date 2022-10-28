package util

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {
	asserts := assert.New(t)

	asserts.True(Exist("C:\\Windows"))
	asserts.False(Exist("C:\\Windows\\not_exist"))
}

func TestGzipCompress(t *testing.T) {
	asserts := assert.New(t)

	buf := bytes.NewBufferString("hello world")
	reader, err := GzipCompress(buf)
	asserts.NoError(err)
	asserts.NotNil(reader)
}
