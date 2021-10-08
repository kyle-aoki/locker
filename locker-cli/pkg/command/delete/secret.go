package delete

import (
	"lkcli/pkg/lpath"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

func DeleteSecret(path string) {
	repoName, envName, secretName := lpath.Split3(path)

	payload := payload.GetPathPayload(repoName, envName, secretName)

	res := request.Post("/secret/delete", payload, true)

	response.PrintMessageResponse(res)
}
