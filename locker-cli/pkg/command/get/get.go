package get

import (
	"fmt"
	"lkcli/pkg/help"
	"lkcli/pkg/logger"
	"lkcli/pkg/lpath"
	"lkcli/pkg/message"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

func GetEnvsOrSecrets(paths ...string) {
	for i, path := range paths {
		executeGetSecrets(path)
		printNewLine(i, paths)
	}
}

func printNewLine(i int, paths []string) {
	if i != len(paths) - 1 {
		fmt.Println()
	}
}

var repo string

func executeGetSecrets(path string) {
	components := lpath.Split(path)

	switch len(components) {
	case 2:
		repo = components[0]
		GetSecrets(repo, components[1])
	case 1:
		if repo == "" {
			help.PrintHelpCommandThenExit()
		}
		GetSecrets(repo, components[0])
	default:
		logger.Exit(message.GetHelp1)
	}
}

func GetSecrets(repoName string, envName string) {
	payload := payload.GetPathPayload(repoName, envName)

	res := request.Post("/secret/get", payload, true)
	response.PrintKeyValueResponse(res)
}
