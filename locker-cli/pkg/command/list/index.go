package list

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/help"
)

func List(args []string) {
	cmd, args := argument.Pop(args)

	switch cmd {
	case "repos":
		switch len(args) {
		case 0:
			listRepos("-1", "-1")
		case 1:
			listRepos(args[0], "-1")
		default:
			listRepos(args[0], args[1])
		}
	case "envs":
		argument.NCheck(1, args)
		getEnvs(args)
	case "secrets":
		arg, _ := argument.Pop(args)
		listSecrets(arg)
	default:
		help.PrintHelpCommandThenExit()
	}
}
