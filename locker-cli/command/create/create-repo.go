package create

import (
	"lkcli/request"
	"lkcli/response"
)

type CreateRepoPayload struct {
	RepoName string `json:"repoName"`
}

func CreateRepo(repoName string) {
	createRepoPayload := CreateRepoPayload{
		RepoName: repoName,
	}

	res := request.Post("/repo/create", createRepoPayload, true)

	response.PrintMessageResponse(res)
}
