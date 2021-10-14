package create

import (
	"lkcli/pkg/argument"
)

func Create(args []string) {
	cmd, args := argument.Pop(args)

	switch cmd {
	case "repo":
		argument.NCheck(1, args)
		createRepo(args...)
	case "env":
		argument.NCheck(1, args)
		createEnv(args...)
	case "secret":
		argument.NCheck(2, args)
		createSecret(args[0], args[1], args[2:]...)
	}
}
