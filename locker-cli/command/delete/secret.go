package delete

import (
	"lkcli/lpath"
	"lkcli/payload"
	"lkcli/request"
	"lkcli/response"
)

func DeleteSecret(path string) {
	repoName, envName, secretName := lpath.Split3(path)

	payload := payload.GetPathPayload(repoName, envName, secretName)

	res := request.Post("/secret/delete", payload, true)

	response.PrintMessageResponse(res)
}
