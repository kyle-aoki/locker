package copy

import "lkcli/pkg/help"

func Copy(args []string) {
	if len(args) == 0 {
		help.PrintHelpCommandThenExit()
	}

	arg, args := args[0], args[1:]
	if !(len(args) >= 2) {
		help.PrintHelpCommandThenExit()
	}

	switch arg {
	case "env":
		path, args := args[0], args[1:]
		CopyEnv(path, args...)
	case "repo":
		path, args := args[0], args[1:]
		CopyRepo(path, args...)
	default:
		help.PrintHelpCommandThenExit()
	}
}
