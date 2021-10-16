package logger

import (
	"fmt"
	"os"
)

func InfoThenExit(message string) {
	fmt.Println(message)
	os.Exit(0)
}

func Exit(message string) {
	fmt.Print(message)
	os.Exit(1)
}
