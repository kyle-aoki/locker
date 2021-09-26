package create

import (
	"lkcli/help"
	"lkcli/arguments")

func Create(args []string) {
	arguments.CheckArgsLength(args, 2, "")

	switch secondCommand := args[1]; secondCommand {
	case "repo":
		CreateRepo(args)
	case "env":
		CreateEnv(args)
	case "secret":
		CreateSecret(args)
	default:
		help.PrintHelpThenExit()
	}
}
