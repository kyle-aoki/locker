package create

import (
	"lkcli/path"
	"lkcli/request"
	"lkcli/response"
)

type CreateRepoPayload struct {
	RepoName string `json:"repoName"`
}

func CreateRepo(args []string) {
	repoName := path.GetRepoName(args)

	createRepoPayload := CreateRepoPayload{
		RepoName: repoName,
	}

	res := request.Post("/repo/create", createRepoPayload, true)

	response.PrintResponseMessage(res)
}
