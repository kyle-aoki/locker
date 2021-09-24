package session

func GetCredentialPaths() (string, string) {
	lockerDir := getLockerDir()
	return lockerDir + slash() + "username.txt", lockerDir + slash() + "session-token.txt"
}
