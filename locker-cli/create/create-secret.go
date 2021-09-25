package create

import (
	"lkcli/path"
	"lkcli/req"
	"lkcli/response"
	"log"
	"strings"
)

type CreateSecretPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
	SecretName string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

func CreateSecret(args []string) {
	if len(args) < 4 {
		log.Fatal("Try: lk create secret <repo>/<env>/<secret-name> <secret-value>")
	}

	repoName, envName, secretName := path.GetSecretPathComponents(args)
	secretValue := strings.Join(args[3:], " ")

	createSecretPayload := CreateSecretPayload{
		RepoName: repoName,
		EnvName: envName,
		SecretName: secretName,
		SecretValue: secretValue,
	}

	res := req.Post("/secret/create", createSecretPayload, true)

	response.PrintResponseMessage(res)
}


