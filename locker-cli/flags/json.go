package flags

import "os"

func JsonFlag() bool {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "--json" {
			return true
		}
	}
	return false
}
