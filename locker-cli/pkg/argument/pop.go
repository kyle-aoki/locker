package argument

func Pop(args []string) (string, []string) {
	ZeroCheck(args)
	return args[0], args[1:]
}

func Pop2(args []string) (string, string, []string) {
	NCheck(2, args)
	return args[0], args[1], args[2:]
}
