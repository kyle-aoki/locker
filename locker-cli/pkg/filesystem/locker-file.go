package filesystem

import (
	"fmt"
	"lkcli/pkg/filesystem/util"
	"lkcli/pkg/logger"
	"lkcli/pkg/message"
)

func GetHostFilePath() string {
	return GetLockerDir() + util.Slash() + "host.txt"
}

func GetHost() string {
	host, err := util.ReadFile(GetHostFilePath())
	if err != nil {
		logger.Exit(message.Host1)
	}
	return host
}

func PrintHost() {
	hostName := GetHost()
	fmt.Print(hostName)
}

func GetUsernameFilePath() string {
	return GetLockerDir() + util.Slash() + "username.txt"
}

func GetUsername() string {
	username, err := util.ReadFile(GetUsernameFilePath())
	if err != nil {
		logger.Exit(message.LogIn1)
	}
	return username
}

func GetSessionTokenFilePath() string {
	return GetLockerDir() + util.Slash() + "session-token.txt"
}

func GetSessionToken() string {
	sessionToken, err := util.ReadFile(GetSessionTokenFilePath())
	if err != nil {
		logger.Exit(message.LogIn1)
	}
	return sessionToken
}

func GetCredentials() (string, string) {
	username := GetUsername()
	sessionToken := GetSessionToken()

	return username, sessionToken
}
