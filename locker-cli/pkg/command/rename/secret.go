package rename

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type RenameSecretPayload struct {
	payload.RepoEnvSecretPayload
	NewSecretName string `json:"newSecretName"`
}

func renameSecret(path string, newSecretName string) {
	repoName, envName, secretName := lpath.Split3(path)
	payload := RenameSecretPayload{
		RepoEnvSecretPayload: payload.RepoEnvSecretPayload{ 
			RepoName: repoName,
			EnvName: envName,
			SecretName: secretName,
		},
		NewSecretName: newSecretName,
	}
	res := request.Post("/secret/rename", payload, true)
	response.PrintMessageResponse(res)
}
