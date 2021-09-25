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

	fmt.Print(mr.message())
}

func PrintResponseSecret(res []byte) {
	sr := SecretResponse{}
	getResponse(res, &sr)

	if sr.ok() {
		fmt.Print(sr.secretValue())
	} else {
		logger.Exit(sr.message())
	}
}

func PrintListResponse(res []byte) {
	lr := ListResponse{}
	getResponse(res, &lr)
	checkAuthError(lr)

	for i, repo := range lr.list() {
		if i == len(lr.list())-1 {
			fmt.Print(repo)
		} else {
			fmt.Println(repo)
		}
	}
}

func checkAuthError(responseClass ResponseClass) {
	if responseClass.code() == "UEAUTH" {
		logger.Exit("Log in first. Try: lk login <username> <password>")
	}
}
