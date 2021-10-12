package get

import (
	"lkcli/pkg/help"
)

func Get(args []string) {
	switch len(args) {
	case 0:
		help.PrintHelpCommandThenExit()
	case 1:
		GetEnvsOrSecrets(args[0])
	default:
		arg, args := args[0], args[1:]
		GetEnvsOrSecrets(arg, args...)
	}
}
