package copy

import (
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type CopyRepoPayload struct {
	RepoName  string `json:"repoName"`
	NewRepoName   string `json:"newRepoName"`
}

func CopyRepo(repoName string, newRepoName string) {

	payload := CopyRepoPayload{
		RepoName:  repoName,
		NewRepoName:   newRepoName,
	}
	res := request.Post("/repo/copy", payload, true)
	response.PrintMessageResponse(res)
}
