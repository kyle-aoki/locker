package delete

import (
	"lkcli/pkg/flags"
	"lkcli/pkg/help"
	"lkcli/pkg/lpath"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type DeleteEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName  string `json:"envName"`
	Force    bool   `json:"force"`
}

var deleteEnvRepo string

func deleteEnv(paths []string) {
	for i, path := range paths {
		if i == 0 {
			repoName, envName := lpath.Split2(path)
			deleteEnvRepo = repoName
			executeDeleteEnv(repoName, envName)
		} else {
			components := lpath.Split(path)
			switch len(components) {
			case 1:
				executeDeleteEnv(deleteEnvRepo, components[0])
			case 2:
				executeDeleteEnv(components[0], components[1])
			default:
				help.PrintHelpCommandThenExit()
			}
		}
	}
}

func executeDeleteEnv(repoName string, envName string) {
	deleteEnvPayload := DeleteEnvPayload{
		RepoName: repoName,
		EnvName:  envName,
		Force:    flags.LockerFlags["--force"],
	}

	res := request.Post("/env/delete", deleteEnvPayload, true)

	response.PrintMessageResponse(res)
}