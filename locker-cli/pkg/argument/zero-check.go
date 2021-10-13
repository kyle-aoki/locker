package argument

import "lkcli/pkg/help"

func ZeroCheck(args []string) {
	if len(args) == 0 {
		help.PrintHelpCommandThenExit()
	}
}

func NCheck(n int, args []string) {
	if !(len(args) >= n) {
		help.PrintHelpCommandThenExit()
	}
}
