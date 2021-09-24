package main

import (
	"lkcli/create"
	"lkcli/help"
	"lkcli/login"
	"lkcli/random"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) == 1 {
		help.PrintHelp()
		return
	}
	args := os.Args[1:]

	switch firstCommand := args[0]; firstCommand {
	case "create":
		create.Create(args)
	case "login":
		login.LogIn(args)
	case "rand":
		random.GenerateRandomString(args)
	default:
		help.PrintHelp()
	}
}
