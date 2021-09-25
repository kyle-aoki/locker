package create

import (
	"lkcli/logger"
	"lkcli/message"
	"lkcli/path"
	"lkcli/req"
	"lkcli/response"
	"strings"
)

type CreateSecretPayload struct {
	RepoName    string `json:"repoName"`
	EnvName     string `json:"envName"`
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

func CreateSecret(args []string) {
	if len(args) < 4 {
		logger.Fatal(message.CreateSecret1)
	}

	repoName, envName, secretName := path.GetSecretPathComponents(args)
	secretValue := strings.Join(args[3:], " ")

	createSecretPayload := CreateSecretPayload{
		RepoName:    repoName,
		EnvName:     envName,
		SecretName:  secretName,
		SecretValue: secretValue,
	}

	res := req.Post("/secret/create", createSecretPayload, true)

	response.PrintResponseMessage(res)
}
