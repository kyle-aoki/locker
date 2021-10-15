package copy

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type CopyEnvPayload struct {
	RepoName       string `json:"repoName"`
	EnvName        string `json:"envName"`
	TargetRepoName string `json:"targetRepoName"`
	TargetEnvName  string `json:"targetEnvName"`
}

func CopyEnv(path string, targetPaths ...string) {
	repoName, envName := lpath.Split2(path)
	for _, targetPath := range targetPaths {
		components := lpath.Split(targetPath)
		if len(components) == 1 {
			ExecuteCopyEnv(repoName, envName, repoName, components[0])
		} else if len(components) == 2 {
			ExecuteCopyEnv(repoName, envName, components[0], components[1])
		}
	}
}

func ExecuteCopyEnv(repoName string, envName string, targetRepoName string, targetEnvName string) {
	payload := CopyEnvPayload{
		RepoName:       repoName,
		EnvName:        envName,
		TargetRepoName: targetRepoName,
		TargetEnvName:  targetEnvName,
	}
	res := request.Post("/env/copy", payload, true)
	response.PrintMessageResponse(res)
}
