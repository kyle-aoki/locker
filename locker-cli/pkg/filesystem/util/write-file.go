package util

import (
	"io/ioutil"
	"lkcli/pkg/logger"
	"os"
)

func WriteFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		logger.Exit("Failed to create " + path)
	}
}
