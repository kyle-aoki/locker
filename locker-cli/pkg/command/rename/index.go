package rename

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/help"
)

func Rename(args []string) {
	cmd, path, newName, args := argument.Pop3(args)

	switch cmd {
	case "repo":
		renameRepo(path, newName)
	case "env":
		renameEnv(path, newName)
	case "secret":
		renameSecret(path, newName)
	default:
		help.PrintHelpCommandThenExit()
	}
}
