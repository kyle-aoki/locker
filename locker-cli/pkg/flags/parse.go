package flags

import (
	"lkcli/pkg/help"
	"lkcli/pkg/version"
	"os"
)

var LockerFlags = map[string]bool {
	"--json": false,
	"--force": false,
	"--help": false,
	"--version": false,
	"-v": false,
	"-h": false,
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

	if LockerFlags["--help"] || LockerFlags["-h"] {
		help.PrintHelpThenExit()
	}
	if LockerFlags["--version"] || LockerFlags["-v"] {
		version.PrintVersionThenExit()
	}

	return args
}

func isFlag(arg string) bool {
	if (arg[0] == '-' && arg[1] == '-') || (arg[0] == '-' && len(arg) == 2) {
		return true
	}
	return false
}

func remove(args []string, i int) []string {
	return append(args[:i], args[i+1:]...)
}
