package help

import (
	"lkcli/logger"
)

const HelpText string = `
	help text goes here
`

func PrintHelp() {
	logger.Exit(HelpText)
}
