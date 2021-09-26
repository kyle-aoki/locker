package response

import (
	"encoding/json"
	"fmt"
	"lkcli/logger"
)

func getResponse(res []byte, responseClass ResponseClass) {
	json.Unmarshal(res, responseClass)
	checkAuthError(responseClass)
}

func PrintResponseMessage(res []byte) {
	mr := MessageResponse{}
	getResponse(res, &mr)

	fmt.Print(mr.Message)
}

func PrintResponseSecret(res []byte) {
	sr := SecretResponse{}
	getResponse(res, &sr)

	if sr.Ok {
		fmt.Print(sr.SecretValue)
	} else {
		logger.Exit(sr.Message)
	}
}

func PrintListResponse(res []byte) {
	lr := ListResponse{}
	getResponse(res, &lr)
	checkAuthError(lr)

	if !lr.Ok {
		logger.Exit(lr.Message)
	}

	for i, repo := range lr.List {
		if i == len(lr.List)-1 {
			fmt.Print(repo)
		} else {
			fmt.Println(repo)
		}
	}
}

func checkAuthError(responseClass ResponseClass) {
	if responseClass.getCode() == "UEAUTH" {
		logger.Exit("Log in first. Try: lk login <username> <password>")
	}
}
