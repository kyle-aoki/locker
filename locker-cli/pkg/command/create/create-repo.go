package create

import (
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type CreateRepoPayload struct {
	RepoName string `json:"repoName"`
}

func createRepo(repoNames ...string) {
	for _, repoName := range repoNames {
		executeCreateRepo(repoName)
	}
}

func executeCreateRepo(repoName string) {
	createRepoPayload := CreateRepoPayload{
		RepoName: repoName,
	}

	res := request.Post("/repo/create", createRepoPayload, true)

	response.PrintMessageResponse(res)
}
