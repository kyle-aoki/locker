package rename

import (
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type RenameRepoPayload struct {
	RepoName    string `json:"repoName"`
	NewRepoName string `json:"newRepoName"`
}

func renameRepo(repoName string, newRepoName string) {
	renameRepoPayload := RenameRepoPayload{
		RepoName:    repoName,
		NewRepoName: newRepoName,
	}

	res := request.Post("/repo/rename", renameRepoPayload, true)

	response.PrintMessageResponse(res)
}
