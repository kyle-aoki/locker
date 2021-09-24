package create

import (
	"encoding/json"
	"fmt"
	"lkcli/req"
)

type CreateRepoPayload struct {
	RepoName string `json:"repoName"`
}

type CreateRepoResponse struct {
	Message string `json:"message"`
}

func CreateRepo(args []string) {
	path := args[2]
	createRepoPayload := CreateRepoPayload{
		RepoName: path,
	}

	res := req.Post("/repo/create", createRepoPayload, true)

	createRepoResponse := CreateRepoResponse{}
	json.Unmarshal([]byte(res), &createRepoResponse)

	fmt.Print(createRepoResponse.Message)
}
