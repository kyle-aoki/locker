package create

import (
	"lkcli/request"
	"lkcli/response"
	"lkcli/arguments")

type CreateEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName  string `json:"envName"`
}

func CreateEnv(args []string) {
	repoName, envName := arguments.GetEnvPathComponents(args)

	createEnvPayload := CreateEnvPayload{
		RepoName: repoName,
		EnvName:  envName,
	}

	res := request.Post("/env/create", createEnvPayload, true)

	response.PrintResponseMessage(res)
}
