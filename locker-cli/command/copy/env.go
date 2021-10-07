package copy

import (
	"lkcli/lpath"
	"lkcli/request"
	"lkcli/response"
)

type CopyEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
	TargetEnv string `json:"targetEnv"`
}

func CopyEnv(path string, targetEnv string) {
	repoName, envName := lpath.Split2(path)

	payload := CopyEnvPayload{
		RepoName: repoName,
		EnvName: envName,
		TargetEnv: targetEnv,
	}
	res := request.Post("/env/copy", payload, true)
	response.PrintMessageResponse(res)
}
