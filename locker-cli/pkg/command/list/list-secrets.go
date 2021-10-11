package list

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type ListSecretsPayload struct {
	RepoName string `json:"repoName"`
	EnvName  string `json:"envName"`
}

func listSecrets(path string) {
	repoName, envName := lpath.Split2(path)

	listSecretsPayload := ListSecretsPayload{
		RepoName: repoName,
		EnvName:  envName,
	}

	res := request.Post("/secret/list", listSecretsPayload, true)
	response.PrintListResponse(res)
}
