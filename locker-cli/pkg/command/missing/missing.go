package missing

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type MissingSecretsRequest struct {
	RepoName      string `json:"repoName"`
	EnvName       string `json:"envName"`
	TargetEnvName string `json:"targetEnvName"`
}

func Missinge(path string, targetEnv string) {
	repoName, envName := lpath.Split2(path)

	payload := MissingSecretsRequest{
		RepoName:      repoName,
		EnvName:       envName,
		TargetEnvName: targetEnv,
	}

	res := request.Post("/env/missing", payload, true)
	response.PrintListResponse(res)
}
