package util

import (
	"errors"
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("")
	}
	return string(bytes), nil
}
