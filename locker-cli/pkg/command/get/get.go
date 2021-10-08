package get

import (
	"lkcli/pkg/logger"
	"lkcli/pkg/lpath"
	"lkcli/pkg/message"
	"lkcli/pkg/payload"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

func Get(path string) {
	components := lpath.Split(path)

	if len(components) > 3 || len(components) < 1 {
		logger.Exit(message.GetHelp1)
	}

	payload := payload.GetPathPayload(components...)

	res := request.Post("/env/get", payload, true)
	response.PrintKeyValueResponse(res)
}
