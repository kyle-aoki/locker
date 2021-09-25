package help

import "lkcli/logger"

func CheckLength(args []string, length int, message string) {
	if len(args) < length {
		if len(message) > 0 {
			logger.Exit(message)
		} else {
			PrintHelpThenExit()
		}
	}
}
