package set

import (
	"lkcli/filesystem"
	"lkcli/help"
	"lkcli/logger"
	"lkcli/message"
)

func Set(args []string) {
	help.CheckArgsLength(args, 2, message.Host2)

	switch command := args[1]; command {
	case "host":
		filesystem.SaveHost(args)
	default:
		logger.Exit(message.Host2)
	}
}
