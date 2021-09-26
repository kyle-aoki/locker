package create

import (
	"lkcli/logger"
	"lkcli/message"
	"lkcli/request"
	"lkcli/response"
	"strings"
	"lkcli/arguments")

type CreateSecretPayload struct {
	RepoName    string `json:"repoName"`
	EnvName     string `json:"envName"`
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

func CreateSecret(args []string) {
	if len(args) < 4 {
		logger.Exit(message.CreateSecret1)
	}

	repoName, envName, secretName := arguments.GetSecretPathComponents(args, 2)
	secretValue := strings.Join(args[3:], " ")

	createSecretPayload := CreateSecretPayload{
		RepoName:    repoName,
		EnvName:     envName,
		SecretName:  secretName,
		SecretValue: secretValue,
	}

	res := request.Post("/secret/create", createSecretPayload, true)

	response.PrintResponseMessage(res)
}
