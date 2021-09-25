package response

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintResponseMessage(res string) {
	lockerServerResponse := LockerServerResponse{}
	json.Unmarshal([]byte(res), &lockerServerResponse)

	fmt.Print(lockerServerResponse.Message)
}

func PrintResponseSecret(res string) {
	lockerServerResponse := LockerServerResponse{}
	json.Unmarshal([]byte(res), &lockerServerResponse)

	if (lockerServerResponse.Ok) {
		fmt.Print(lockerServerResponse.SecretValue)
	} else {
		log.Fatal(lockerServerResponse.Message)
	}
}
