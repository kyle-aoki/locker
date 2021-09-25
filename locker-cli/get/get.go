package get

import (
	"lkcli/path"
	"lkcli/request"
	"lkcli/response"
)

type GetSecretPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
	SecretName string `json:"secretName"`
}

func Get(args []string) {
	repoName, envName, secretName := path.GetSecretPathComponentsFrom1(args)

	getSecretPayload := GetSecretPayload{
		RepoName: repoName,
		EnvName: envName,
		SecretName: secretName,
	}

	res := request.Post("/secret/get", getSecretPayload, true)

	response.PrintResponseSecret(res)
}
