package create

import "lkcli/pkg/help"

func Create(args []string) {
	if len(args) == 0 {
		help.PrintHelpCommandThenExit()
	}
	switch arg, args := args[0], args[1:]; arg {
	case "repo":
		createRepo(args...)
	case "env":
	case "secret":
	}
}
