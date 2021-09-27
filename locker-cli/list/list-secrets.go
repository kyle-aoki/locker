package list

import (
	"lkcli/arguments"
	"lkcli/request"
	"lkcli/response"
)

type ListSecretsPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
}

func ListSecrets(args []string) {
	arguments.CheckArgsLength(args, 3, "")

	repoName, envName := arguments.GetEnvPathComponents(args, 2)

	listSecretsPayload := ListSecretsPayload{
		RepoName: repoName,
		EnvName: envName,
	}

	res := request.Post("/secret/list", listSecretsPayload, true)
	response.PrintListResponse(res)
}
