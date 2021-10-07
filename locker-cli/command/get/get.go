package get

import (
	"lkcli/logger"
	"lkcli/lpath"
	"lkcli/message"
	"lkcli/payload"
	"lkcli/request"
	"lkcli/response"
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
