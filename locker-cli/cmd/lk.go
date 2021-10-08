package main

import (
	"lkcli/pkg/argument"
	"lkcli/pkg/command/copy"
	"lkcli/pkg/command/create"
	"lkcli/pkg/command/delete"
	"lkcli/pkg/command/get"
	"lkcli/pkg/command/list"
	"lkcli/pkg/command/login"
	"lkcli/pkg/command/missing"
	"lkcli/pkg/command/rand"
	"lkcli/pkg/command/set"
	"lkcli/pkg/command/update"
	"lkcli/pkg/flags"
	"lkcli/pkg/help"
)

func main() {
	Args := flags.Parse()
	argument.ParseArguments(Args)

	switch argument.ArgLen {
	case 0:
		help.PrintHelpThenExit()
	case 1:
		help.PrintHelpThenExit()
	case 2: // ----------------------------------------------------------------------------------
		switch argument.FirstArgument {
		case "list":
			switch argument.SecondArgument {
			case "repos":
				list.ListRepos(argument.Args2, "-1", "-1")
			}
		case "get":
			get.Get(argument.SecondArgument)
		case "rand":
			rand.GenerateRandomString(argument.SecondArgument)
		case "set":
			set.SetHost(argument.SecondArgument)
		default:
			help.PrintHelpThenExit()
		}
	case 3: // ----------------------------------------------------------------------------------
		switch argument.FirstArgument {
		case "copy":
			copy.CopyEnv(argument.SecondArgument, argument.ThirdArgument)
		case "delete":
			switch argument.SecondArgument {
			case "secret":
				delete.DeleteSecret(argument.ThirdArgument)
			case "env":
				delete.DeleteEnv(argument.ThirdArgument)
			}
		case "list":
			switch argument.SecondArgument {
			case "repos":
				list.ListRepos(argument.Args3, argument.ThirdArgument, "-1")
			case "secrets":
				list.ListSecrets(argument.ThirdArgument)
			}
		case "create":
			switch argument.SecondArgument {
			case "repo":
				create.CreateRepo(argument.ThirdArgument)
			case "env":
				create.CreateEnv(argument.ThirdArgument)
			}
		case "missing":
			missing.Missing(argument.SecondArgument, argument.ThirdArgument)
		case "login":
			login.LogIn(argument.SecondArgument, argument.ThirdArgument)
		default:
			help.PrintHelpThenExit()
		}
	case 4: // ----------------------------------------------------------------------------------
		switch argument.FirstArgument {
		case "list":
			switch argument.SecondArgument {
			case "repos":
				list.ListRepos(argument.Args3, argument.ThirdArgument, argument.FourthArgument)
			}
		case "create":
			switch argument.SecondArgument {
			case "secret":
				create.CreateSecret(argument.ThirdArgument, argument.FourthArgument, false)
			}
		case "update":
			switch argument.SecondArgument {
			case "repo":
				update.UpdateRepo(argument.ThirdArgument, argument.FourthArgument)
			case "secret":
				update.UpdateSecret(argument.ThirdArgument, argument.FourthArgument, false)
			}
		default:
			help.PrintHelpThenExit()
		}
	default: // ----------------------------------------------------------------------------------
		switch argument.FirstArgument {
		case "create":
			switch argument.SecondArgument {
			case "secret":
				create.CreateSecret(argument.ThirdArgument, argument.FourthArgument, true)
			}
		case "update":
			switch argument.SecondArgument {
			case "secret":
				update.UpdateSecret(argument.ThirdArgument, argument.FourthArgument, true)
			}
		default:
			help.PrintHelpThenExit()
		}
	}
}
