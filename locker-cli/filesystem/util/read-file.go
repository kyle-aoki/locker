package util

import (
	"io/ioutil"
	"lkcli/logger"
)

func ReadFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Exit("Failed to read username file at " + path)
	}
	return string(bytes)
}
