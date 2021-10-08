package response

type OkField struct {
	Ok bool `json:"ok"`
}

type CodeField struct {
	Code string `json:"code"`
}

type MessageField struct {
	Message string `json:"message"`
}

type StrField struct {
	Str string `json:"str"`
}

type ListField struct {
	List []string `json:"list"`
}

type KeyValue struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type KeyValueListField struct {
	KeyValues []KeyValue `json:"keyValueList"`
}
