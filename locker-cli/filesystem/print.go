package filesystem

import (
	"fmt"
)

func PrintHost() {
	hostName := GetHost()
	fmt.Print(hostName)
}
