package get

import (
	"lkcli/request"
	"lkcli/response"
	"lkcli/arguments")

type GetSecretPayload struct {
	RepoName   string `json:"repoName"`
	EnvName    string `json:"envName"`
	SecretName string `json:"secretName"`
}

func Get(args []string) {
	repoName, envName, secretName := arguments.GetSecretPathComponents(args, 1)

	getSecretPayload := GetSecretPayload{
		RepoName:   repoName,
		EnvName:    envName,
		SecretName: secretName,
	}

	res := request.Post("/secret/get", getSecretPayload, true)

	response.PrintResponseSecret(res)
}
