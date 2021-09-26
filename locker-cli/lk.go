package main

import (
	"lkcli/create"
	"lkcli/filesystem"
	"lkcli/get"
	"lkcli/help"
	"lkcli/list"
	"lkcli/login"
	"lkcli/random"
	"lkcli/set"
	"lkcli/update"
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
	case "list":
		list.List(args)
	case "host":
		filesystem.PrintHost()
	case "update":
		update.Update(args)
	default:
		help.PrintHelpThenExit()
	}
}
