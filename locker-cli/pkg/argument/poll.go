package argument

func Poll(args []string) (string, []string) {
	ZeroCheck(args)
	return args[0], args[1:]
}

func Poll2(args []string) (string, string, []string) {
	NCheck(2, args)
	return args[0], args[1], args[2:]
}

func Poll3(args []string) (string, string, string, []string) {
	NCheck(3, args)
	return args[0], args[1], args[2], args[3:]
}
