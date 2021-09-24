package create

import "lkcli/help"

func Create(args []string) {
	switch secondCommand := args[1]; secondCommand {
	case "repo":
		CreateRepo(args)
	case "env":
		CreateEnv(args)
	case "secret":
		CreateSecret(args)
	default:
		help.PrintHelp()
	}
}
