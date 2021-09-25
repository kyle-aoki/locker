package login

import (
	"encoding/json"
	"fmt"
	"lkcli/req"
	"lkcli/filesystem"
	"log"
)

type LogInPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInResponse struct {
	Ok bool `json:"ok"`
	Message string `json:"message"`
	SessionToken string `json:"sessionToken"`
}

func LogIn(args []string) {
	if len(args) < 3 {
		log.Fatal("Try: lk login <username> <password>")
	}
	username := args[1]
	password := args[2]

	logInPayload := LogInPayload{
		Username: username,
		Password: password,
	}

	res := req.Post("/user/public/log-in", logInPayload, false)

	logInResponse := LogInResponse{}
	json.Unmarshal([]byte(res), &logInResponse)

	if !logInResponse.Ok {
		log.Fatal(logInResponse.Message)
	}

	filesystem.SaveUsernameAndSessionToken(username, logInResponse.SessionToken)
	fmt.Print(logInResponse.Message)
}
