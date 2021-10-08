package set

import (
	"lkcli/pkg/filesystem"
	"lkcli/pkg/filesystem/util"
	"lkcli/pkg/logger"
)

func SetHost(host string) {
	filesystem.InitLockerDir()

	hostFilePath := filesystem.GetHostFilePath()

	util.WriteFile(hostFilePath, host)

	logger.Exit("Set host to: " + host)
}
