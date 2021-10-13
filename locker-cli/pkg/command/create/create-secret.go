package create

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
	"strings"
)

type CreateSecretPayload struct {
	RepoName    string `json:"repoName"`
	EnvName     string `json:"envName"`
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

func createSecret(path string, secretValue string, additionalValues ...string) {
	repoName, envName, secretName := lpath.Split3(path)

	finalSecretValue := secretValue + strings.Join(additionalValues, " ")

	createSecretPayload := CreateSecretPayload{
		RepoName:    repoName,
		EnvName:     envName,
		SecretName:  secretName,
		SecretValue: finalSecretValue,
	}

	res := request.Post("/secret/create", createSecretPayload, true)

	response.PrintMessageResponse(res)
}
