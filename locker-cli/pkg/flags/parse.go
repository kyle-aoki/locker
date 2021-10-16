package flags

import (
	"os"
)

var LockerFlags = map[string]bool{
	"--json":  false,
	"--force": false,

	"--help": false, "-h": false,

	"--version": false, "-v": false,
}

func Parse() []string {
	args := os.Args[1:]

	var argsRemoved int
	for i, arg := range args {
		if isFlag(arg) {
			LockerFlags[arg] = true
			args = poll(args, i-argsRemoved)
			argsRemoved += 1
		}
	}

	CheckImmediateExitFlags()

	return args
}

func isFlag(arg string) bool {
	if (arg[0] == '-' && arg[1] == '-') || (arg[0] == '-' && len(arg) == 2) {
		return true
	}
	return false
}

func poll(args []string, i int) []string {
	return append(args[:i], args[i+1:]...)
}
