package create

import (
	"lkcli/argument"
	"lkcli/lpath"
	"lkcli/request"
	"lkcli/response"
)

type CreateSecretPayload struct {
	RepoName    string `json:"repoName"`
	EnvName     string `json:"envName"`
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

func CreateSecret(path string, secretValue string, moreThanFourArgs bool) {
	repoName, envName, secretName := lpath.Split3(path)

	if moreThanFourArgs {
		secretValue += argument.VarArguments
	}

	createSecretPayload := CreateSecretPayload{
		RepoName:    repoName,
		EnvName:     envName,
		SecretName:  secretName,
		SecretValue: secretValue,
	}

	res := request.Post("/secret/create", createSecretPayload, true)

	response.PrintMessageResponse(res)
}
