package response

type BasicResponse struct {
	OkField
	CodeField
}

type MessageResponse struct {
	BasicResponse
	MessageField
}

type StrResponse struct {
	BasicResponse
	StrField
}

type ListResponse struct {
	BasicResponse
	ListField
}

type KeyValueResponse struct {
	BasicResponse
	KeyValueListField
}

type ErrorResponse struct {
	BasicResponse
	MessageResponse
}
