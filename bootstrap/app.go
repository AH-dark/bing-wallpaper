package bootstrap

import (
	_ "embed"
	"fmt"
)

//go:embed cover
var text string

func InitApplication(path string) {
	fmt.Println(text)

	Init(path)
}
