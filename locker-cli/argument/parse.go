package argument

import (
	"fmt"
	"lkcli/help"
	"os"
)

var ArgLen int
var FirstArgument string
var SecondArgument string
var ThirdArgument string
var FourthArgument string
var VarArguments string

func ParseArguments() {
	args := os.Args[1:]

	if len(args) == 0 {
		help.PrintHelpThenExit()
	}
	ArgLen = len(args)

	for i := range args {
		GetArgumente(args, i)
	}
}

func GetArgumente(args []string, argIndex int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	switch argIndex {
	case 0:
		FirstArgument = Get(args, argIndex)
	case 1:
		SecondArgument = Get(args, argIndex)
	case 2:
		ThirdArgument = Get(args, argIndex)
	case 3:
		FourthArgument = Get(args, argIndex)
	default:
		VarArguments += Get(args, argIndex)
	}
}

func Get(args []string, argIndex int) string {
	return args[argIndex]
}
