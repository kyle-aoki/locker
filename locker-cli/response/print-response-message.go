package response

import (
	"encoding/json"
	"fmt"
	"lkcli/logger"
)

func getLockerResponse(res string) LockerServerResponse {
	lockerServerResponse := LockerServerResponse{}
	json.Unmarshal([]byte(res), &lockerServerResponse)

	checkAuthError(lockerServerResponse)

	return lockerServerResponse
}

func PrintResponseMessage(res string) {
	lockerServerResponse := getLockerResponse(res)
	fmt.Print(lockerServerResponse.Message)
}

func PrintResponseSecret(res string) {
	lockerServerResponse := getLockerResponse(res)
	if lockerServerResponse.Ok {
		fmt.Print(lockerServerResponse.SecretValue)
	} else {
		logger.Exit(lockerServerResponse.Message)
	}
}

func checkAuthError(lockerServerResponse LockerServerResponse) {
	if lockerServerResponse.Code == "UEAUTH" {
		logger.Exit("Log in first. Try: lk login <username> <password>")
	}
}
