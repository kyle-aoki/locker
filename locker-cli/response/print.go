package response

import (
	"encoding/json"
	"fmt"
	"lkcli/logger"
)

func CheckAuthError(r HasCode) {
	if code := r.getCode(); code == "UEAUTH" {
		fmt.Println(r)
		logger.Exit("Please log in first.")
	}
}

func PrintMessageResponse(res []byte) {
	r := MessageResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	fmt.Print(r.Message)
}

func PrintListResponse(res []byte) {
	r := ListResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	fmt.Print(r.List)
}

func PrintStrResponse(res []byte) {
	r := StrResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	fmt.Print(r.Str)
}

func PrintKeyValueResponse(res []byte) {
	r := KeyValueResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	fmt.Println(r)
	fmt.Print(r.KeyValues)
}
