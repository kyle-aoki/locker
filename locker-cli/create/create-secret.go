package create

import (
	"encoding/json"
	"fmt"
	"lkcli/req"
	"log"
	"strings"
)

type CreateSecretPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
	SecretName string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

type CreateSecretResponse struct {
	Message string `json:"message"`
}

func CreateSecret(args []string) {
	path := args[2]
	repoName, envName, secretName := getSecretPathComponents(path)
	secretValue := args[3]

	createSecretPayload := CreateSecretPayload{
		RepoName: repoName,
		EnvName: envName,
		SecretName: secretName,
		SecretValue: secretValue,
	}

	res := req.Post("/secret/create", createSecretPayload, true)

	createSecretResponse := CreateSecretResponse{}
	json.Unmarshal([]byte(res), &createSecretResponse)

	fmt.Print(createSecretResponse.Message)
}

func getSecretPathComponents(path string) (string, string, string) {
	components := strings.Split(path, "/")
	if len(components) != 3 {
		log.Fatal("Invalid path.")
	}

	return components[0], components[1], components[2]
}
