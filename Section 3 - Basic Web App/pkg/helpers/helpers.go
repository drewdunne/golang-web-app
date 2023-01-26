package helpers

import (
	"fmt"
	"os"
)

func Pwd() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return path
}