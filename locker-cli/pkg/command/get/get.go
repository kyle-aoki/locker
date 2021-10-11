package get

import (
	"lkcli/pkg/logger"
	"lkcli/pkg/lpath"
	"lkcli/pkg/message"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

func Gete(path string) {
	components := lpath.Split(path)

	switch len(components) {
	case 2:
		GetSecrets(components)
	case 1:
		GetEnvs(components)
	default:
		logger.Exit(message.GetHelp1)
	}
}

func GetEnvs(components []string) {
	payload := payload.GetPathPayload(components...)

	res := request.Post("/repo/get", payload, true)
	response.PrintListResponse(res)
}

func GetSecrets(components []string) {
	payload := payload.GetPathPayload(components...)

	res := request.Post("/env/get", payload, true)
	response.PrintKeyValueResponse(res)
}
