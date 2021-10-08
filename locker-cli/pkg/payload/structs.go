package payload

type RepoPayload struct {
	RepoName string `json:"repoName"`
}
type RepoEnvPayload struct {
	RepoName string `json:"repoName"`
	EnvName  string `json:"envName"`
}
type RepoEnvSecretPayload struct {
	RepoName   string `json:"repoName"`
	EnvName    string `json:"envName"`
	SecretName string `json:"secretName"`
}
type RepoEnvSecretValuePayload struct {
	RepoName    string `json:"repoName"`
	EnvName     string `json:"envName"`
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}
