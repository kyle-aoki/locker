package copy

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type CopyEnvPayload struct {
	RepoName  string `json:"repoName"`
	EnvName   string `json:"envName"`
	TargetRepoName string `json:"targetRepoName"`
	TargetEnvName string `json:"targetEnvName"`
}

func CopyEnv(path string, targetPaths ...string) {
	for _, targetPath := range targetPaths {
		ExecuteCopyEnv(path, targetPath)
	}
}

func ExecuteCopyEnv(path string, targetPath string) {
	repoName, envName := lpath.Split2(path)
	targetRepoName, targetEnvName := lpath.Split2(targetPath)

	payload := CopyEnvPayload{
		RepoName:  repoName,
		EnvName:   envName,
		TargetRepoName: targetRepoName,
		TargetEnvName: targetEnvName,
	}
	res := request.Post("/env/copy", payload, true)
	response.PrintMessageResponse(res)
}
