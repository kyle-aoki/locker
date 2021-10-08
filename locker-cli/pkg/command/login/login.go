package login

import (
	"encoding/json"
	"fmt"
	"lkcli/pkg/filesystem"
	"lkcli/pkg/request"
	"lkcli/pkg/response"
)

type LogInPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LogIn(username string, password string) {
	logInPayload := LogInPayload{
		Username: username,
		Password: password,
	}

	res := request.Post("/user/public/log-in", logInPayload, false)

	r := response.StrResponse{}
	json.Unmarshal([]byte(res), &r)

	filesystem.SaveUsernameAndSessionToken(username, r.Str)
	fmt.Println("Successfully logged in.")
}
