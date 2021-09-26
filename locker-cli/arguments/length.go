package arguments

import (
	"lkcli/help"
	"lkcli/logger"
)

func CheckArgsLength(args []string, length int, message string) {
	if len(args) < length {
		if len(message) > 0 {
			logger.Exit(message)
		} else {
			help.PrintHelpThenExit()
		}
	}
}
