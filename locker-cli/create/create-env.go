package create

import (
	"lkcli/request"
	"lkcli/response"
	"lkcli/path"
)

type CreateEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
}

func CreateEnv(args []string) {
	repoName, envName := path.GetEnvPathComponents(args)

	createEnvPayload := CreateEnvPayload{
		RepoName: repoName,
		EnvName: envName,
	}

	res := request.Post("/env/create", createEnvPayload, true)

	response.PrintResponseMessage(res)
}
