package response

// Components
type OkStruct struct {
	Ok bool `json:"ok"`
}
type CodeStruct struct {
	Code string `json:"code"`
}
type MessageStruct struct {
	Message string `json:"message"`
}
type SecretValueStruct struct {
	SecretValue string `json:"secretValue"`
}
type ListStruct struct {
	List []string `json:"list"`
}

// Interfaces
type ResponseClass interface {
	ok()
	getCode() string
}
type MessageClass interface {
	message()
}
type SecretValueClass interface {
	secretValue()
}
type ListClass interface {
	list()
}

type Response struct {
	OkStruct
	CodeStruct
}

func (r Response) ok() {}
func (r Response) getCode() string {
	return r.Code
}

type MessageResponse struct {
	Response
	MessageStruct
}

func (mr MessageResponse) message() {}

type SecretResponse struct {
	MessageResponse
	SecretValueStruct
}

func (sr SecretResponse) secretValue() {}

type ListResponse struct {
	Response
	MessageStruct
	ListStruct
}

func (lr ListResponse) list() {}


