package version

import "lkcli/pkg/logger"

const version string = "1.0.0"

func PrintVersionThenExit() {
	logger.Exit(version)
}
