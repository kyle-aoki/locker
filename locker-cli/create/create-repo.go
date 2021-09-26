package create

import (
	"lkcli/request"
	"lkcli/response"
	"lkcli/arguments"
)

type CreateRepoPayload struct {
	RepoName string `json:"repoName"`
}

func CreateRepo(args []string) {
	repoName := arguments.GetRepoName(args)

	createRepoPayload := CreateRepoPayload{
		RepoName: repoName,
	}

	res := request.Post("/repo/create", createRepoPayload, true)

	response.PrintResponseMessage(res)
}
