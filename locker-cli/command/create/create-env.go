package create

import (
	"lkcli/lpath"
	"lkcli/request"
	"lkcli/response"
)

type CreateEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName  string `json:"envName"`
}

func CreateEnv(path string) {
	repoName, envName := lpath.Split2(path)

	createEnvPayload := CreateEnvPayload{
		RepoName: repoName,
		EnvName:  envName,
	}

	res := request.Post("/env/create", createEnvPayload, true)

	response.PrintMessageResponse(res)
}
