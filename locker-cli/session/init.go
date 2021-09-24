package session

import "os"

func initStashLocation() {
	etcDir := getEtcDir()
	etcExists, _ := exists(etcDir)

	if !etcExists {
		createEtcDir()
	}

	secretsDir := getSecretsDir()
	secretsDirExists, _ := exists(secretsDir)

	if !secretsDirExists {
		createSecretsDir()
	}

	lockerDir := getLockerDir()
	lockerDirExists, _ := exists(lockerDir)

	if !lockerDirExists {
		createLockerDir()
	}
}

func getLockerDir() string {
	secretsDir := getSecretsDir()
	lockerDir := secretsDir + slash() + "locker"
	return lockerDir
}

func createLockerDir() {
	os.MkdirAll(getLockerDir(), os.ModePerm)
}

func getSecretsDir() string {
	etcDir := getEtcDir()
	secretsDir := etcDir + slash() + "secrets"
	return secretsDir
}

func createSecretsDir() {
	os.MkdirAll(getSecretsDir(), os.ModePerm)
}

func getEtcDir() string {
	root := getOsRootPath()
	etcDir := root + slash() + "etc"
	return etcDir
}

func createEtcDir() {
	os.MkdirAll(getEtcDir(), os.ModePerm)
}