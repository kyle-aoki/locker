package get

import (
	"lkcli/arguments"
	"lkcli/logger"
	"lkcli/request"
	"lkcli/response"
)

type GetSecretPayload struct {
	RepoName   string `json:"repoName"`
	EnvName    string `json:"envName"`
	SecretName string `json:"secretName"`
}

type GetEnvPayload struct {
	RepoName   string `json:"repoName"`
	EnvName    string `json:"envName"`
}

type GetRepoPayload struct {
	RepoName   string `json:"repoName"`
}

func Get(args []string) {
	components := arguments.GetVariableSecretPathComponents(args, 1)

	switch len(components) {
	case 3:
		GetSecret(components)
	case 2:
		GetAllSecrets(components)
	case 1:
		GetAllEnvs(components)
	default:
		logger.Exit("Try: lk get <repo>\nor\nlk get <repo>/<env>\nor\nlk get <repo>/<env>/<secret>")
	}
}

func GetAllEnvs(components []string) {
	payload := GetRepoPayload{
		RepoName:   components[0],
	}

	res := request.Post("/repo/get", payload, true)
	response.PrintResponseSecret(res)
}

func GetAllSecrets(components []string) {
	payload := GetEnvPayload{
		RepoName:   components[0],
		EnvName:    components[1],
	}

	res := request.Post("/env/get", payload, true)

	response.PrintRawSecrerts(res)
}

func GetSecret(components []string) {
	payload := GetSecretPayload{
		RepoName:   components[0],
		EnvName:    components[1],
		SecretName: components[2],
	}

	res := request.Post("/secret/get", payload, true)
	response.PrintResponseSecret(res)
}
