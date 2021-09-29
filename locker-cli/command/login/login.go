package login

import (
	"encoding/json"
	"fmt"
	"lkcli/filesystem"
	"lkcli/logger"
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

func LogIn(username string, password string) {
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
