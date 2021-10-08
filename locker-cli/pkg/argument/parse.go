package argument

import (
	"lkcli/pkg/help"
)

var ArgLen int
var FirstArgument string
var SecondArgument string
var ThirdArgument string
var FourthArgument string
var VarArguments string

func ParseArguments(args []string) {
	if len(args) == 0 {
		help.PrintHelpThenExit()
	}
	ArgLen = len(args)

	for i := range args {
		GetArgument(args, i)
	}
}

func GetArgument(args []string, argIndex int) {
	switch argIndex {
	case 0:
		FirstArgument = args[argIndex]
	case 1:
		SecondArgument = args[argIndex]
	case 2:
		ThirdArgument = args[argIndex]
	case 3:
		FourthArgument = args[argIndex]
	case 4:
		VarArguments += args[argIndex]
	default:
		VarArguments += " " + args[argIndex]
	}
}
