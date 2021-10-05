package payload

func GetPathPayload(names ...string) interface{} {
	switch len(names) {
	case 0:
		panic("Illegal state exception in GetPathPayload.")
	case 1:
		return getRepoPayload(names[0])
	case 2:
		return getRepoEnvPayload(names[0], names[1])
	case 3:
		return getRepoEnvSecretPayload(names[0], names[1], names[2])
	case 4:
		return getRepoEnvSecretValuePayload(names[0], names[1], names[2], names[3])
	default:
		panic("Illegal state exception in GetPathPayload.")
	}
}

func getRepoPayload(repoName string) RepoPayload {
	return RepoPayload{
		RepoName: repoName,
	}
}

func getRepoEnvPayload(repoName string, envName string) RepoEnvPayload {
	return RepoEnvPayload{
		RepoName: repoName,
		EnvName:  envName,
	}
}

func getRepoEnvSecretPayload(repoName string, envName string, secretName string) RepoEnvSecretPayload {
	return RepoEnvSecretPayload{
		RepoName:   repoName,
		EnvName:    envName,
		SecretName: secretName,
	}
}

func getRepoEnvSecretValuePayload(
	repoName string,
	envName string,
	secretName string,
	secretValue string,
) RepoEnvSecretValuePayload {
	return RepoEnvSecretValuePayload{
		RepoName:   repoName,
		EnvName:    envName,
		SecretName: secretName,
		SecretValue: secretValue,
	}
}
