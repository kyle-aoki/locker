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
	ok() bool
	code() string
}
type MessageClass interface {
	message() string
}
type SecretValueClass interface {
	secretValue() string
}
type ListClass interface {
	list() []string
}

// Structs
type Response struct {
	OkStruct
	CodeStruct
}
func (r Response) ok() bool {
	return r.Ok
}
func (r Response) code() string {
	return r.Code
}

type MessageResponse struct {
	Response
	MessageStruct
}
func (mr MessageResponse) ok() bool {
	return mr.Ok
}
func (mr MessageResponse) code() string {
	return mr.Code
}
func (mr MessageResponse) message() string {
	return mr.Message
}

type SecretResponse struct {
	Response
	MessageStruct
	SecretValueStruct
}
func (sr SecretResponse) 	ok() bool{
	return sr.Ok
}
func (sr SecretResponse) code() string {
	return sr.Code
}
func (sr SecretResponse) message() string{
	return sr.Message
}
func (sr SecretResponse) 	secretValue() string{
	return sr.SecretValue
}

type ListResponse struct {
	Response
	ListStruct
}
func (lr ListResponse) ok() bool {
	return lr.Ok
}
func (lr ListResponse) code() string {
	return lr.Code
}
func (lr ListResponse) list() []string {
	return lr.List
}
