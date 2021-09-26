package update

import (
	"lkcli/arguments"
	"lkcli/request"
	"lkcli/response"
)

type UpdateSecretPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
	SecretName string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

func updateSecret(args []string) {
	arguments.CheckArgsLength(args, 4, "Try: lk update secret <repo>/<env>/<secret> <secret-value>")

	repoName, envName, secretName := arguments.GetSecretPathComponents(args, 2)
	newSecretValue := args[3]

	updateSecretPayload := UpdateSecretPayload{
		RepoName: repoName,
		EnvName: envName,
		SecretName: secretName,
		SecretValue: newSecretValue,
	}

	res := request.Post("/secret/update", updateSecretPayload, true)

	response.PrintResponseMessage(res)
}
