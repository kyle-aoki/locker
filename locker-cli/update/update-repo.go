package update

import (
	"lkcli/arguments"
	"lkcli/request"
	"lkcli/response"
)

type UpdateRepoPayload struct {
	RepoName string `json:"repoName"`
	NewRepoName string `json:"newRepoName"`
}

func updateRepo(args []string) {
	arguments.CheckArgsLength(args, 4, "Try: lk update repo <repo> <new-repo-name>")

	repoName := args[2]
	newRepoName := args[3]

	updateRepoPayload := UpdateRepoPayload{
		RepoName: repoName,
		NewRepoName: newRepoName,
	}

	res := request.Post("/repo/update-name", updateRepoPayload, true)

	response.PrintResponseMessage(res)
}