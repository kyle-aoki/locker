package list

import (
	"lkcli/help"
	"lkcli/path"
)

func List(args []string) {
	path.CheckArgsLength(args, 2, help.ListHelpText)

	switch listType := args[1]; listType {
	case "repos":
		ListRepos(args)
	}
}
