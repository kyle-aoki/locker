package delete

import (
	"lkcli/flags"
	"lkcli/lpath"
	"lkcli/request"
	"lkcli/response"
)

type DeleteEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
	Force bool `json:"force"`
}

func DeleteEnv(path string) {
	repoName, envName := lpath.Split2(path)

	deleteEnvPayload := DeleteEnvPayload{
		RepoName: repoName,
		EnvName: envName,
		Force: flags.LockerFlags["--force"],
	}

	res := request.Post("/env/delete", deleteEnvPayload, true)

	response.PrintMessageResponse(res)
}
