package list

import (
	"lkcli/pkg/help"
	"lkcli/pkg/logger"
	"lkcli/pkg/message"
)

func List(args []string) {
	if len(args) == 0 {
		help.PrintHelpCommandThenExit()
	}
	arg, args := args[0], args[1:]

	switch arg {
	case "repos":
		switch len(args) {
		case 0:
			listRepos("-1", "-1")
		case 1:
			listRepos(args[0], "-1")
		default:
			listRepos(args[0], args[1])
		}
	case "secrets":
		if len(args) == 0 {
			logger.Exit(message.ListSecrets1)
		}
		listSecrets(args[0])
	default:
		help.PrintHelpCommandThenExit()
	}
}
