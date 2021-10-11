package update

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type UpdateSecretPayload struct {
	RepoName    string `json:"repoName"`
	EnvName     string `json:"envName"`
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

func UpdateSecret(path string, newSecretValue string, moreThanFourArgs bool) {
	repoName, envName, secretName := lpath.Split3(path)

	updateSecretPayload := UpdateSecretPayload{
		RepoName:    repoName,
		EnvName:     envName,
		SecretName:  secretName,
		SecretValue: newSecretValue,
	}

	res := request.Post("/secret/update", updateSecretPayload, true)

	response.PrintMessageResponse(res)
}
