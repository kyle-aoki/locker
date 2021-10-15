package create

import (
	"lkcli/pkg/help"
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type CreateEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName  string `json:"envName"`
}

var createEnvRepo string

func createEnv(paths ...string) {
	for i, path := range paths {
		if i == 0 {
			repoName, envName := lpath.Split2(path)
			createEnvRepo = repoName
			executeCreateEnv(repoName, envName)
		} else {
			components := lpath.Split(path)
			switch len(components) {
			case 1:
				executeCreateEnv(createEnvRepo, path)
			case 2:
				executeCreateEnv(components[0], components[1])
			default:
				help.PrintHelpCommandThenExit()
			}
		}
	}
}

func executeCreateEnv(repoName string, envName string) {
	createEnvPayload := CreateEnvPayload{
		RepoName: repoName,
		EnvName:  envName,
	}

	res := request.Post("/env/create", createEnvPayload, true)

	response.PrintMessageResponse(res)
}
