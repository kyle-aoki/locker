package rand

import "lkcli/pkg/argument"

func Rand(args []string) {
	argument.ZeroCheck(args)
	generateRandomString(args[0])
}
