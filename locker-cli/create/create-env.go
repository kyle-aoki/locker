package create

import (
	"encoding/json"
	"fmt"
	"lkcli/req"
	"log"
	"strings"
)

type CreateEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
}

type CreateEnvResponse struct {
	Message string `json:"message"`
}

func CreateEnv(args []string) {
	path := args[2]
	repoName, envName := getEnvPathComponents(path)

	createEnvPayload := CreateEnvPayload{
		RepoName: repoName,
		EnvName: envName,
	}

	res := req.Post("/env/create", createEnvPayload, true)

	createEnvResponse := CreateEnvResponse{}
	json.Unmarshal([]byte(res), &createEnvResponse)

	fmt.Print(createEnvResponse.Message)
}

func getEnvPathComponents(path string) (string, string) {
	components := strings.Split(path, "/")
	if len(components) != 2 {
		log.Fatal("Invalid path.")
	}
	return components[0], components[1]
}
