package list

import (
	"fmt"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

func getEnvs(args []string) {
	for i, repoName := range args {
		executeGetEnvs(repoName)
		if i != len(args) - 1 {
			fmt.Println()
		}
	}
}

func executeGetEnvs(repoName string) {
	payload := payload.GetPathPayload(repoName)

	res := request.Post("/env/get", payload, true)
	response.PrintListResponse(res)
}