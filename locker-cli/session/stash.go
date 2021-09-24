package session

import (
	"io/ioutil"
	"log"
	"os"
)

func Stash(username string, sessionToken string) {
	initStashLocation()

	usernameFilePath := getLockerDir() + slash() + "username.txt"
	sessionTokenFilePath := getLockerDir() + slash() + "session-token.txt"

	err := ioutil.WriteFile(usernameFilePath, []byte(username), os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create " + usernameFilePath)
	}

	err = ioutil.WriteFile(sessionTokenFilePath, []byte(sessionToken), os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create " + sessionTokenFilePath)
	}
}
