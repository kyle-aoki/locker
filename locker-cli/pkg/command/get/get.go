package get

import (
	"fmt"
	"lkcli/pkg/logger"
	"lkcli/pkg/lpath"
	"lkcli/pkg/message"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

func GetEnvsOrSecrets(path string, additionalEnvs ...string) {
	components := lpath.Split(path)

	switch len(components) {
	case 2:
		repoName, envName := components[0], components[1]
		if len(additionalEnvs) == 0 {
			GetSecrets(repoName, envName)
		} else {
			fmt.Println(envName)
			GetSecrets(repoName, envName)
			for _, env := range additionalEnvs {
				fmt.Println(env)
				GetSecrets(repoName, env)
			}
		}
	case 1:
		GetEnvs(components)
	default:
		logger.Exit(message.GetHelp1)
	}
}

func GetEnvs(components []string) {
	payload := payload.GetPathPayload(components...)

	res := request.Post("/env/get", payload, true)
	response.PrintListResponse(res)
}

func GetSecrets(repoName string, envName string) {
	payload := payload.GetPathPayload(repoName, envName)

	res := request.Post("/secret/get", payload, true)
	response.PrintKeyValueResponse(res)
}
