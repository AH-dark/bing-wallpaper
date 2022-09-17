package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvStr(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("default", EnvStr("not_exist", "default"))

	os.Setenv("exist", "value")
	asserts.Equal("value", EnvStr("exist", "default"))

	os.Unsetenv("exist")
	asserts.Equal("default", EnvStr("exist", "default"))
}

func TestEnvInt(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal(1, EnvInt("not_exist", 1))

	os.Setenv("exist", "2")
	asserts.Equal(2, EnvInt("exist", 1))

	os.Unsetenv("exist")
	asserts.Equal(1, EnvInt("exist", 1))

	os.Setenv("exist", "not_int")
	asserts.Equal(1, EnvInt("exist", 1))
}

func TestEnvArr(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal([]string{"default"}, EnvArr("not_exist", []string{"default"}))

	os.Setenv("exist", "value")
	asserts.Equal([]string{"value"}, EnvArr("exist", []string{"default"}))

	os.Unsetenv("exist")
	asserts.Equal([]string{"default"}, EnvArr("exist", []string{"default"}))

	os.Setenv("exist", "value1,value2")
	asserts.Equal([]string{"value1", "value2"}, EnvArr("exist", []string{"default"}))

	os.Unsetenv("exist")
	asserts.Equal([]string{"default"}, EnvArr("exist", []string{"default"}))

	os.Setenv("exist", "value1,value2,value3")
	asserts.Equal([]string{"value1", "value2", "value3"}, EnvArr("exist", []string{"default"}))

	os.Unsetenv("exist")
	asserts.Equal([]string{"default"}, EnvArr("exist", []string{"default"}))
}
