package set

import (
	"lkcli/filesystem"
	"lkcli/help"
	"lkcli/logger"
	"lkcli/message"
)

func Set(args []string) {
	help.CheckLength(args, 2, message.SetHelp1)

	switch command := args[1]; command {
	case "host":
		filesystem.SaveHost(args)
	default:
		logger.Exit(message.SetHelp1)
	}
}
