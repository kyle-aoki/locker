package response

type LockerServerResponse struct {
	Ok bool `json:"ok"`
	Code bool `json:"code"`
	Message string `json:"message"`
	SecretValue string `json:"secretValue"`
}
