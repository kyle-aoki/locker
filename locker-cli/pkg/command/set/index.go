package set

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/help"
)

func Set(args []string) {
	target, value, args := argument.Pop2(args)
	switch target {
	case "host":
		setHost(value)
	default:
		help.PrintHelpCommandThenExit()
	}
}
