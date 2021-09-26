package help

import (
	"lkcli/logger"
)

const ListHelpText string = `
Commands:

lk list repos <limit?> <offset?>
`

func PrintListHelpThenExit() {
	logger.Exit(ListHelpText)
}