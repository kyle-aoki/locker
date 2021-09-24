package session

func GetCredentialPaths() (string, string) {
	lockerDir := getLockerDir()
	
	usernamePath := lockerDir + slash() + "username.txt"
	sessionTokenPath := lockerDir + slash() + "session-token.txt"

	return usernamePath, sessionTokenPath
}
