package session

import "runtime"

func getOsRootPath() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "C:\\"
	default:
		return "/"
	}
}

func slash() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "\\"
	default:
		return "/"
	}
}
