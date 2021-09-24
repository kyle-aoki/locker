package login

import (
	"encoding/json"
	"fmt"
	"lkcli/req"
	"lkcli/session"
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

	session.Stash(username, logInResponse.SessionToken)

	fmt.Println(logInResponse.Message)
}
