package util

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read username file at " + path)
	}
	return string(bytes)
}
