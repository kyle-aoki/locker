package filesystem

import (
	"lkcli/filesystem/util"
	"lkcli/logger"
)

func SaveHost(args []string) {
	if len(args) < 2 {
		logger.Exit("Try: lk set host https://host-name.com")
	}

	host := args[2]

	InitLockerDir()

	hostFilePath := GetHostFilePath()

	util.WriteFile(hostFilePath, host)

	logger.Exit("Set host to: " + host)
}
