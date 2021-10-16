package host

import (
	"lkcli/pkg/filesystem"
	"lkcli/pkg/logger"
)

func getHost() {
	logger.InfoThenExit(filesystem.GetHost()) 
}
