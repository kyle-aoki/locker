package lpath

import (
	"lkcli/pkg/logger"
	"strings"
)

func Split(s string) []string {
	return strings.Split(s, "/")
}

func Split2(s string) (string, string) {
	components := strings.Split(s, "/")
	if len(components) != 2 {
		logger.Exit("Invalid path. Try: <repo>/<env>")
	}
	return components[0], components[1]
}

func Split3(s string) (string, string, string) {
	components := strings.Split(s, "/")
	if len(components) != 3 {
		logger.Exit("Invalid path. Try: <repo>/<env>/<secret>")
	}
	return components[0], components[1], components[2]
}
