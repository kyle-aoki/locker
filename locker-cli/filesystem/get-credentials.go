package filesystem

import (
	"lkcli/filesystem/util"
	"lkcli/logger"
	"lkcli/message"
)

func GetCredentials() (string, string) {
	usernamePath := GetUsernameFilePath()
	sessionTokenPath := GetSessionTokenFilePath()

	usernameFileExists, _ := util.Exists(usernamePath)
	sessionTokenFileExists, _ := util.Exists(sessionTokenPath)

	if !usernameFileExists || !sessionTokenFileExists {
		logger.Exit(message.LogIn1)
	}

	username := util.ReadFile(usernamePath)
	sessionToken := util.ReadFile(sessionTokenPath)

	return username, sessionToken
}
