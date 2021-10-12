package login

import "lkcli/pkg/help"

func Login(args []string) {
	if !(len(args) >= 2) {
		help.PrintHelpCommandThenExit()
	}
	LogIn(args[0], args[1])
}
