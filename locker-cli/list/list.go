package list

import "lkcli/path"

func List(args []string) {
	path.CheckArgsLength(args, 2, "")

	switch listType := args[1]; listType {
	case "repos":
		ListRepos()
	}
}
