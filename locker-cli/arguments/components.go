package arguments

import (
	"lkcli/logger"
	"lkcli/message"
	"strings"
)

func GetRepoName(args []string) string {
	CheckArgsLength(args, 3, message.CreateRepo1)
	return args[2]
}

func GetEnvPathComponents(args []string, argIndex int) (string, string) {
	CheckArgsLength(args, argIndex + 1, message.CreateEnv1)

	path := args[argIndex]

	components := strings.Split(path, "/")
	if len(components) != 2 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1]
}

func GetSecretPathComponents(args []string, argIndex int) (string, string, string) {
	CheckArgsLength(args, argIndex, "")

	path := args[argIndex]

	components := strings.Split(path, "/")

	if len(components) != 3 {
		logger.Exit("Invalid path.")
	}

	return components[0], components[1], components[2]
}

func GetVariableSecretPathComponents(args []string, argIndex int) []string {
	CheckArgsLength(args, argIndex, "")

	path := args[argIndex]

	components := strings.Split(path, "/")

	return components
}
