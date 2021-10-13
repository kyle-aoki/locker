package create

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type CreateEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName  string `json:"envName"`
}

func createEnv(paths ...string) {
	for _, path := range paths {
		executeCreateEnv(path)
	}
}

func executeCreateEnv(path string) {
	repoName, envName := lpath.Split2(path)

	createEnvPayload := CreateEnvPayload{
		RepoName: repoName,
		EnvName:  envName,
	}

	res := request.Post("/env/create", createEnvPayload, true)

	response.PrintMessageResponse(res)
}
