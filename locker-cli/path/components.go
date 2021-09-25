package path

import (
	"lkcli/logger"
	"strings"
)

func GetRepoName(args []string) string {
	return args[2]
}

func GetEnvPathComponents(args []string) (string, string) {
	path := args[2]

	components := strings.Split(path, "/")
	if len(components) != 2 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1]
}

func GetSecretPathComponents(args []string) (string, string, string) {
	path := args[2]

	components := strings.Split(path, "/")

	if len(components) != 3 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1], components[2]
}

func GetSecretPathComponentsFrom1(args []string) (string, string, string) {
	if len(args) < 2 {
		logger.Exit("Try: lk get <repo>/<env>/<secret-name>")
	}

	path := args[1]
	components := strings.Split(path, "/")

	if len(components) != 3 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1], components[2]
}
