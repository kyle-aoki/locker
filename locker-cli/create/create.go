package create

func Create(args []string) {
	secondCommand := args[1]
	
	if (secondCommand == "repo") {
		CreateRepo(args)
	}
}
