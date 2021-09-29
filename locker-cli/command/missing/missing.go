package missing

import (
	"lkcli/lpath"
	"lkcli/request"
	"lkcli/response"
)

type MissingSecretsRequest struct {
	RepoName string `json:"repoName"`
	EnvName string `json:"envName"`
	TargetEnvName string `json:"targetEnvName"`
}

func Missing(path string, targetEnv string) {
	repoName, envName := lpath.Split2(path)

	payload := MissingSecretsRequest{
		RepoName: repoName,
		EnvName: envName,
		TargetEnvName: targetEnv,
	}

	res := request.Post("/env/missing", payload, true)
	response.PrintListResponse(res)
}
