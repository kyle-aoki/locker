package logger

import (
	"fmt"
	"os"
)

func Info(message string) {
	fmt.Println(message)
}

func Exit(message string) {
	fmt.Print(message)
	os.Exit(0)
}
