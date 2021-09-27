package update

import (
	"lkcli/arguments"
	"lkcli/help"
)

func Update(args []string) {
	arguments.CheckArgsLength(args, 2, "")

	switch secondCommand := args[1]; secondCommand {
	case "secret":
		updateSecret(args)
	case "repo":
		updateRepo(args)
	default:
		help.PrintHelpThenExit()
	}
}
