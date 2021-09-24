package main

import (
	"lkcli/create"
	"lkcli/help"
	"lkcli/login"
	"os"
)

func main() {
	args := os.Args[1:]

	switch firstCommand := args[0]; firstCommand {
	case "create":
		create.Create(args)
	case "login":
		login.LogIn(args)
	default:
		help.PrintHelp()
	}
}
