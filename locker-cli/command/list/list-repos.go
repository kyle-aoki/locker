package list

import (
	"lkcli/argument"
	"lkcli/logger"
	"lkcli/message"
	"lkcli/request"
	"lkcli/response"
	"strconv"
)

type ListReposPayload struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

func ListRepos(aq argument.ArgQuantity, limitStr string, offsetStr string) {
	limit, offset := ToInt(limitStr, offsetStr)

	listReposPayload := ListReposPayload{
		Limit: limit,
		Offset: offset,
	}

	res := request.Post("/repo/list", listReposPayload, true)
	response.PrintListResponse(res)
}

func getLimitAndOffset(args []string) (int, int) {
	if len(args) == 2 {
		return -1, -1
	}
	if len(args) == 3 {
		limit, err := strconv.ParseInt(args[2], 10, 32)
		if err != nil {
			logger.Exit(message.ListRepos1)
		}
		return int(limit), -1
	}
	if len(args) <= 4 {
		limit, err := strconv.ParseInt(args[2], 10, 32)
		if err != nil {
			logger.Exit(message.ListRepos1)
		}
		offset, err := strconv.ParseInt(args[3], 10, 32)
		if err != nil {
			logger.Exit(message.ListRepos1)
		}
		return int(limit), int(offset)
	}
	return -1, -1
}

func ToInt(s1 string, s2 string) (int, int) {
	n1, _ := strconv.ParseInt(s1, 10, 32)
	n2, _ := strconv.ParseInt(s2, 10, 32)
	return int(n1), int(n2)
}
