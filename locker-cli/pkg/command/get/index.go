package get

import (
	"lkcli/pkg/argument"
)

func Get(args []string) {
	argument.ZeroCheck(args)
	GetEnvsOrSecrets(args[0])
}
