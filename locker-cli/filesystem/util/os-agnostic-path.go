package util

import "runtime"

func GetOsRootPath() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "C:\\"
	default:
		return "/"
	}
}

func Slash() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "\\"
	default:
		return "/"
	}
}
