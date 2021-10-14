package get

import (
	"lkcli/pkg/argument"
)

func Get(args []string) {
	argument.NCheck(1, args)
	GetEnvsOrSecrets(args...)
}
