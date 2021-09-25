package util

import (
	"io/ioutil"
	"log"
	"os"
)

func WriteFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create " + path)
	}
}
