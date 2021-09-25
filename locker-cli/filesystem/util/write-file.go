package util

import (
	"io/ioutil"
	"lkcli/logger"
	"os"
)

func WriteFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		logger.Fatal("Failed to create " + path)
	}
}
