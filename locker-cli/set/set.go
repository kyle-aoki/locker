package set

import (
	"lkcli/filesystem"
	"lkcli/logger"
	"lkcli/message"
	"lkcli/arguments")

func Set(args []string) {
	arguments.CheckArgsLength(args, 2, message.Host2)

	switch command := args[1]; command {
	case "host":
		filesystem.SaveHost(args)
	default:
		logger.Exit(message.Host2)
	}
}
