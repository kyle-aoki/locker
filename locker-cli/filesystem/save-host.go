package filesystem

import "log"
import "lkcli/filesystem/util"

func SaveHost(args []string) {
	if len(args) < 2 {
		log.Fatal("Try: lk set host https://host-name.com")
	}

	host := args[2]

	InitLockerDir()

	hostFilePath := GetHostFilePath()

	util.WriteFile(hostFilePath, host)
}
