package main

import (
	"lkcli/create"
	"lkcli/filesystem"
	"lkcli/get"
	"lkcli/help"
	"lkcli/login"
	"lkcli/random"
	"lkcli/set"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		help.PrintHelpThenExit()
	}
	args := os.Args[1:]

	switch firstCommand := args[0]; firstCommand {
	case "create":
		create.Create(args)
	case "login":
		login.LogIn(args)
	case "rand":
		random.GenerateRandomString(args)
	case "get":
		get.Get(args)
	case "set":
		set.Set(args)
	case "host":
		filesystem.PrintHost()
	default:
		help.PrintHelpThenExit()
	}
}
