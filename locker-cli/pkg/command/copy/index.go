package copy

import "lkcli/pkg/help"

func Copy(args []string) {
	if len(args) == 0 {
		help.PrintHelpCommandThenExit()
	}

	switch arg, args := args[0], args[1:]; arg {
	case "env":
		if !(len(args) >= 2) {
			help.PrintHelpCommandThenExit()
		}
		path, args := args[0], args[1:]
		CopyEnv(path, args...)
	case "repo":
		if !(len(args) >= 2) {
			help.PrintHelpCommandThenExit()
		}
		path, args := args[0], args[1:]
		CopyRepo(path, args...)
	}
}
