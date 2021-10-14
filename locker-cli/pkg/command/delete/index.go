package delete

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/help"
)

func Delete(args []string) {
	cmd, path, _ := argument.Pop2(args)

	switch cmd {
	case "repo":
		deleteRepo(path)
	case "env":
		deleteEnv(path)
	case "secret":
		deleteSecret(path)
	default:
		help.PrintHelpCommandThenExit()
	}
}
