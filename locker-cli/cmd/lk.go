package main

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/command/copy"
	"lkcli/pkg/command/create"
	"lkcli/pkg/command/delete"
	"lkcli/pkg/command/envs"
	"lkcli/pkg/command/get"
	"lkcli/pkg/command/host"
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
	cmd, args := argument.Pop(args)

	switch cmd {
	case "version": version.PrintVersionThenExit()
	case "help":    help.PrintHelpThenExit()
	case "repos": 	list.List(args)
	case "get":			get.Get(args)
	case "copy":		copy.Copy(args)
	case "create": 	create.Create(args)
	case "delete": 	delete.Delete(args)
	case "login": 	login.Login(args)
	case "set": 		set.Set(args)
	case "update": 	update.Update(args)
	case "rename": 	rename.Rename(args)
	case "rand": 		rand.Rand(args)
	case "missing": missing.Missing(args)
	case "list": 		list.List(args)
	case "envs":    envs.Envs(args)
	case "host":    host.Host()
	default: 				help.PrintHelpCommandThenExit()
	}
}
