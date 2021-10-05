package payload

func GetPathPayload(subpaths ...string) interface{} {
	switch len(subpaths) {
	case 0:
		panic("Illegal state exception in GetPathPayload.")
	case 1:
		return getRepoPayload(subpaths[0])
	case 2:
		return getRepoEnvPayload(subpaths[0], subpaths[1])
	case 3:
		return getRepoEnvSecretPayload(subpaths[0], subpaths[1], subpaths[2])
	case 4:
		return getRepoEnvSecretValuePayload(subpaths[0], subpaths[1], subpaths[2], subpaths[3])
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

func getRepoEnvSecretPayload(
	repoName string,
	envName string,
	secretName string,
) RepoEnvSecretPayload {
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
		RepoName:    repoName,
		EnvName:     envName,
		SecretName:  secretName,
		SecretValue: secretValue,
	}
}
