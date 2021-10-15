package response

import (
	"encoding/json"
	"fmt"
	"lkcli/pkg/logger"
)

func rr(str string) {
	fmt.Println(str)
}

func somethingWentWrong(res []byte) {
	errResp := ErrorResponse{}
	json.Unmarshal(res, &errResp)
	logger.Exit(errResp.Message)
}
