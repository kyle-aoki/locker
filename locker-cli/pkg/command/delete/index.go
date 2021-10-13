package delete

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/help"
)

func Delete(args []string) {
	argument.NCheck(2, args)

	switch arg, args := args[0], args[1:]; arg {
	case "repo":
		deleteRepo(args[0])
	case "env":
		deleteEnv(args[0])
	case "secret":
		deleteSecret(args[0])
	default:
		help.PrintHelpCommandThenExit()
	}
}
