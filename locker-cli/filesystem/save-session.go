package filesystem

import "lkcli/filesystem/util"

func SaveUsernameAndSessionToken(username string, sessionToken string) {
	InitLockerDir()

	usernameFilePath := GetLockerDir() + util.Slash() + "username.txt"
	sessionTokenFilePath := GetLockerDir() + util.Slash() + "session-token.txt"

	util.WriteFile(usernameFilePath, username)
	util.WriteFile(sessionTokenFilePath, sessionToken)
}
