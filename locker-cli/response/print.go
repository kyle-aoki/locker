package response

import (
	"encoding/json"
	"fmt"
	"lkcli/logger"
)

func rr(str string) {
	fmt.Println(str)
}

func CheckAuthError(r HasCode) {
	if code := r.getCode(); code == "UEAUTH" {
		logger.Exit("Please log in first.")
	}
}

func PrintMessageResponse(res []byte) {
	r := MessageResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	rr(r.Message)
}

func PrintListResponse(res []byte) {
	r := ListResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	var printStr string
	for i, str := range r.List {
		printStr += str + addNewLine(len(r.List), i)
	}
	rr(printStr)
}

func addNewLine(length int, index int) string {
	if length - 1 == index {
		return ""
	}
	return "\n"
}

func PrintStrResponse(res []byte) {
	r := StrResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	rr(r.Str)
}
