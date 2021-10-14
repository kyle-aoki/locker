package envs

import (
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

func getEnvs(args []string) {
	for _, repoName := range args {
		executeGetEnvs(repoName)
	}
}

func executeGetEnvs(repoName string) {
	payload := payload.GetPathPayload(repoName)

	res := request.Post("/env/get", payload, true)
	response.PrintListResponse(res)
}
