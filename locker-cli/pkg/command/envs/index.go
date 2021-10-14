package envs

import "lkcli/pkg/argument"

func Envs(args []string) {
	argument.NCheck(1, args)
	getEnvs(args)
}