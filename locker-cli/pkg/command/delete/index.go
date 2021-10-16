package delete

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/help"
)

func Delete(args []string) {
	cmd, args := argument.Poll(args)

	switch cmd {
	case "repo":
		deleteRepo(args)
	case "env":
		deleteEnv(args)
	case "secret":
		deleteSecret(args)
	default:
		help.PrintHelpCommandThenExit()
	}
}
