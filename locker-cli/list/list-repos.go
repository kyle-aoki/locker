package list

import (
	"lkcli/req"
	"lkcli/response"
)

func ListRepos() {
	res := req.Post("/repo/list", nil, true)
	response.PrintListResponse(res)
}
