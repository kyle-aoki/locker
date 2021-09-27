package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lkcli/flags"
	"lkcli/logger"
)

type RawSecretResponse struct {
	Response
	KeyValueList
}

type KeyValueList struct {
	List []KeyValueStruct `json:"list"`
}

type KeyValueStruct struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OutputFormat int
const (
	Ini OutputFormat = iota
	Json
)

func PrintRawSecrerts(res []byte) {
	lr := RawSecretResponse{}
	getResponse(res, &lr)
	checkAuthError(lr)

	outputFormat := Ini
	if flags.JsonFlag() {
		outputFormat = Json
	}

	fullOutputStr := ""
	for i, secret := range lr.List {
		switch outputFormat {
		case Ini:
			fullOutputStr += getIniOutput(secret.Key, secret.Value) + addNewLine(lr, i)
		case Json:
			fullOutputStr += getJsonOutput(secret.Key, secret.Value, isLastInList(lr, i))
		}
	}

	switch outputFormat {
	case Ini:
		fmt.Print(fullOutputStr)
	case Json:
		var out bytes.Buffer
		err := json.Indent(&out, wrapWithBraces(fullOutputStr), "", "    ")
		if err != nil {
			logger.Exit("Error converting to json.")
		}
		fmt.Print(string(out.Bytes()))
	}
}

func isLastInList(lr RawSecretResponse, i int) bool {
	return i == len(lr.List)-1
}

func getIniOutput(key string, value string) string {
	return key + "=" + value
}
func addNewLine(lr RawSecretResponse, i int) string {
	if !isLastInList(lr, i) {
		return "\n"
	}
	return ""
}

func getJsonOutput(key string, value string, last bool) string {
	json := "\"" + key + "\": " + "\"" + value + "\""
	if !last {
		json += ","
	}
	return json
}

func wrapWithBraces(str string) []byte {
	return []byte("{" + str + "}")
}
