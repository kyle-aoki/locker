package response

import (
	"encoding/json"
	"fmt"
	"lkcli/pkg/flags"
)

func PrintKeyValueResponse(res []byte) {
	r := KeyValueResponse{}
	json.Unmarshal(res, &r)
	CheckAuthError(r)

	if !r.Ok {
		somethingWentWrong(res)
	}

	if flags.LockerFlags["--json"] {
		jsonPrint(r)
		return
	}
	iniPrint(r)
}

func iniPrint(r KeyValueResponse) {
	var printStr string
	for i, kv := range r.KeyValues {
		printStr += kv.Key + "=" + kv.Value + addNewLine(len(r.KeyValues), i)
	}
	rr(printStr)
}

func addComma(r KeyValueResponse, i int) string {
	if len(r.KeyValues)-1 == i {
		return ""
	}
	return ","
}

func jsonPrint(r KeyValueResponse) {
	var printStr string = "{\n"
	for i, kv := range r.KeyValues {
		printStr += "\t\"" + kv.Key + "\": \"" + kv.Value + "\"" + addComma(r, i) + "\n"
	}
	printStr += "}\n"
	fmt.Print(printStr)
}
