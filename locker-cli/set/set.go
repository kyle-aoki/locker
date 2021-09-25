package set

import (
	"lkcli/filesystem"
	"lkcli/help"
)

func Set(args []string) {
	switch command := args[1]; command {
	case "host":
		filesystem.SaveHost(args)
	default:
		help.PrintHelp()
	}
}
