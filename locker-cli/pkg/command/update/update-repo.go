package update

import (
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type UpdateRepoPayload struct {
	RepoName    string `json:"repoName"`
	NewRepoName string `json:"newRepoName"`
}

func UpdateRepo(repoName string, newRepoName string) {
	updateRepoPayload := UpdateRepoPayload{
		RepoName:    repoName,
		NewRepoName: newRepoName,
	}

	res := request.Post("/repo/update-name", updateRepoPayload, true)

	response.PrintMessageResponse(res)
}
