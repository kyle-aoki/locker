package filesystem

import "os"
import "lkcli/filesystem/util"

func InitLockerDir() {
	etcDir := getEtcDir()
	etcExists, _ := util.Exists(etcDir)

	if !etcExists {
		createEtcDir()
	}

	secretsDir := getSecretsDir()
	secretsDirExists, _ := util.Exists(secretsDir)

	if !secretsDirExists {
		createSecretsDir()
	}

	lockerDir := GetLockerDir()
	lockerDirExists, _ := util.Exists(lockerDir)

	if !lockerDirExists {
		createLockerDir()
	}
}

func GetLockerDir() string {
	secretsDir := getSecretsDir()
	lockerDir := secretsDir + util.Slash() + "locker"
	return lockerDir
}

func createLockerDir() {
	os.MkdirAll(GetLockerDir(), os.ModePerm)
}

func getSecretsDir() string {
	etcDir := getEtcDir()
	secretsDir := etcDir + util.Slash() + "secrets"
	return secretsDir
}

func createSecretsDir() {
	os.MkdirAll(getSecretsDir(), os.ModePerm)
}

func getEtcDir() string {
	root := util.GetOsRootPath()
	etcDir := root + util.Slash() + "etc"
	return etcDir
}

func createEtcDir() {
	os.MkdirAll(getEtcDir(), os.ModePerm)
}
