package delete

import (
	"lkcli/pkg/flags"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type DeleteRepoPayload struct {
	payload.RepoPayload
	flags.ForceStruct
}

func deleteRepo(args []string) {
	for _, repoName := range args {
		executeDeleteRepo(repoName)
	}
}

func executeDeleteRepo(repoName string) {
	payload := DeleteRepoPayload{
		payload.RepoPayload{
			RepoName: repoName,
		},
		flags.ForceStruct{
			Force: flags.LockerFlags["--force"],
		},
	}

	res := request.Post("/repo/delete", payload, true)
	response.PrintMessageResponse(res)
}
