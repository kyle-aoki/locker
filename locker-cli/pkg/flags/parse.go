package flags

import (
	"os"
)

var LockerFlags = map[string]bool {
	"--json": false,
	"--force": false,
}

func Parse() []string {
	args := os.Args[1:]

	var argsRemoved int
	for i, arg := range args {
		if isFlag(arg) {
			LockerFlags[arg] = true
			args = remove(args, i - argsRemoved)
			argsRemoved += 1
		}
	}
	return args
}

func isFlag(arg string) bool {
	if arg[0] == '-' && arg[1] == '-' {
		return true
	}
	return false
}

func remove(args []string, i int) []string {
	return append(args[:i], args[i+1:]...)
}
