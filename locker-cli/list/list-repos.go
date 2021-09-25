package list

import (
	"lkcli/request"
	"lkcli/response"
)

func ListRepos() {
	res := request.Post("/repo/list", nil, true)
	response.PrintListResponse(res)
}
