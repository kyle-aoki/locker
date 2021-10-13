package main

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/command/copy"
	"lkcli/pkg/command/create"
	"lkcli/pkg/command/delete"
	"lkcli/pkg/command/get"
	"lkcli/pkg/command/list"
	"lkcli/pkg/command/login"
	"lkcli/pkg/command/missing"
	"lkcli/pkg/command/rand"
	"lkcli/pkg/command/rename"
	"lkcli/pkg/command/set"
	"lkcli/pkg/command/update"
	"lkcli/pkg/flags"
	"lkcli/pkg/help"
	"lkcli/pkg/version"
)

func main() {
	args := flags.Parse()

	argument.ZeroCheck(args)

	switch args[0] {
	case "version": version.PrintVersionThenExit()
	case "help":    help.PrintHelpThenExit()
	case "repos": 	list.List(args[1:])
	case "get":			get.Get(args[1:])
	case "copy":		copy.Copy(args[1:])
	case "create": 	create.Create(args[1:])
	case "delete": 	delete.Delete(args[1:])
	case "login": 	login.Login(args[1:])
	case "set": 		set.Set(args[1:])
	case "update": 	update.Update(args[1:])
	case "rename": 	rename.Rename(args[1:])
	case "rand": 		rand.Rand(args[1:])
	case "missing": missing.Missing(args[1:])
	case "list": 		list.List(args[1:])
	default: 				help.PrintHelpCommandThenExit()
	}
}
