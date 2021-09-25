package filesystem

import "lkcli/filesystem/util"

func GetHostFilePath() string {
	return GetLockerDir() + util.Slash() + "host.txt"
}

func GetHost() string {
	return util.ReadFile(GetHostFilePath())
}

func GetUsernameFilePath() string {
	return GetLockerDir() + util.Slash() + "username.txt"
}

func GetUsername() string {
	return util.ReadFile(GetUsernameFilePath())
}

func GetSessionTokenFilePath() string {
	return GetLockerDir() + util.Slash() + "session-token.txt"
}

func GetSessionToken() string {
	return util.ReadFile(GetSessionTokenFilePath())
}
