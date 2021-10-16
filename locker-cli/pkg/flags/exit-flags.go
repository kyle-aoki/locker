package flags

import (
	"lkcli/pkg/help"
	"lkcli/pkg/version"
)

func CheckImmediateExitFlags() {
	if LockerFlags["--help"] || LockerFlags["-h"] {
		help.PrintHelpThenExit()
	}
	if LockerFlags["--version"] || LockerFlags["-v"] {
		version.PrintVersionThenExit()
	}
}
