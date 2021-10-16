package rename

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type RenameEnvPayload struct {
	payload.RepoEnvPayload
	NewEnvName string `json:"newEnvName"`
}

func renameEnv(path string, newEnvName string) {
	repoName, envName := lpath.Split2(path)
	payload := RenameEnvPayload{
		RepoEnvPayload: payload.RepoEnvPayload{
			RepoName: repoName,
			EnvName: envName,
		},
		NewEnvName: newEnvName,
	}
	res := request.Post("/env/rename", payload, true)
	response.PrintMessageResponse(res)
}
