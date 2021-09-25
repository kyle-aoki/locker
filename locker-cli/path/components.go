package path

import (
	"lkcli/logger"
	"lkcli/message"
	"strings"
)

func GetRepoName(args []string) string {
	CheckArgsLength(args, 3, message.CreateRepo1)
	return args[2]
}

func GetEnvPathComponents(args []string) (string, string) {
	CheckArgsLength(args, 3, message.CreateEnv1)

	path := args[2]

	components := strings.Split(path, "/")
	if len(components) != 2 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1]
}

func GetSecretPathComponents(args []string) (string, string, string) {
	CheckArgsLength(args, 3, message.CreateRepo1)

	path := args[2]

	components := strings.Split(path, "/")

	if len(components) != 3 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1], components[2]
}

func GetSecretPathComponentsFrom1(args []string) (string, string, string) {
	CheckArgsLength(args, 2, message.GetSecret1)

	path := args[1]
	components := strings.Split(path, "/")

	if len(components) != 3 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1], components[2]
}
