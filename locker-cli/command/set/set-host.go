package set

import (
	"lkcli/filesystem/util"
	"lkcli/logger"
	"lkcli/filesystem"
)

func SetHost(host string) {
	filesystem.InitLockerDir()

	hostFilePath := filesystem.GetHostFilePath()

	util.WriteFile(hostFilePath, host)

	logger.Exit("Set host to: " + host)
}
