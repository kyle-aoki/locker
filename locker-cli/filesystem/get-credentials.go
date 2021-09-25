package filesystem

import "lkcli/filesystem/util"

func GetCredentials() (string, string) {
	usernamePath := GetUsernameFilePath()
	sessionTokenPath := GetSessionTokenFilePath()

	username := util.ReadFile(usernamePath)
	sessionToken := util.ReadFile(sessionTokenPath)

	return username, sessionToken
}
