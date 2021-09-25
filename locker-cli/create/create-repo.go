package create

import (
	"lkcli/path"
	"lkcli/req"
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

	res := req.Post("/repo/create", createRepoPayload, true)

	response.PrintResponseMessage(res)
}
