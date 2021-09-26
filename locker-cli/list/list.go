package list

import (
	"lkcli/help"
	"lkcli/arguments")

func List(args []string) {
	arguments.CheckArgsLength(args, 2, help.ListHelpText)

	switch listType := args[1]; listType {
	case "repos":
		ListRepos(args)
	}
}
