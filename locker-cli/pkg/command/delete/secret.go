package delete

import (
	"lkcli/pkg/help"
	"lkcli/pkg/lpath"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

var deleteSecretRepo string
var deleteSecretEnv string

func deleteSecret(paths []string) {
	for i, path := range paths {
		if i == 0 {
			repoName, envName, secretName := lpath.Split3(path)
			deleteSecretRepo = repoName
			deleteSecretEnv = envName
			executeDeleteSecret(repoName, envName, secretName)
		} else {
			components := lpath.Split(path)
			switch len(components) {
			case 1:
				executeDeleteSecret(deleteSecretRepo, deleteSecretEnv, components[0])
			case 2:
				executeDeleteSecret(deleteSecretRepo, components[0], components[1])
			case 3:
				executeDeleteSecret(components[0], components[1], components[2])
			default:
				help.PrintHelpCommandThenExit()
			}
		}
	}
}

func executeDeleteSecret(repoName string, envName string, secretName string) {
	payload := payload.GetPathPayload(repoName, envName, secretName)

	res := request.Post("/secret/delete", payload, true)

	response.PrintMessageResponse(res)
}
