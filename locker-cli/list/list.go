package list

import "lkcli/help"

func List(args []string) {
	help.CheckArgsLength(args, 2, "")

	switch listType := args[1]; listType {
	case "repos":
		ListRepos()
	}
}
