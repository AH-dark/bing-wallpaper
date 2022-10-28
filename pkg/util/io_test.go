package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {
	asserts := assert.New(t)

	asserts.True(Exist("C:\\Windows"))
	asserts.False(Exist("C:\\Windows\\not_exist"))
}
