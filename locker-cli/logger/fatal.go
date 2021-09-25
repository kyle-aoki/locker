package logger

import (
	"fmt"
	"os"
)

func Fatal(message string) {
	fmt.Print(message)
	os.Exit(0)
}
