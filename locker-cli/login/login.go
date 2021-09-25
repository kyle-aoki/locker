package login

import (
	"encoding/json"
	"fmt"
	"lkcli/filesystem"
	"lkcli/help"
	"lkcli/logger"
	"lkcli/message"
	"lkcli/request"
)

type LogInPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInResponse struct {
	Ok           bool   `json:"ok"`
	Message      string `json:"message"`
	SessionToken string `json:"sessionToken"`
}

func LogIn(args []string) {
	help.CheckArgsLength(args, 3, message.LogIn2)
	
	username := args[1]
	password := args[2]

	logInPayload := LogInPayload{
		Username: username,
		Password: password,
	}

	res := request.Post("/user/public/log-in", logInPayload, false)

	logInResponse := LogInResponse{}
	json.Unmarshal([]byte(res), &logInResponse)

	if !logInResponse.Ok {
		logger.Exit(logInResponse.Message)
	}

	filesystem.SaveUsernameAndSessionToken(username, logInResponse.SessionToken)
	fmt.Print(logInResponse.Message)
}
