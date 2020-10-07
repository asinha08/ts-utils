package error

import (
	"encoding/json"
	logger "github.com/asinha08/ts-utils/tslogger/fluentd"
	"github.com/asinha08/ts-utils/tsprotos"
	"github.com/golang/protobuf/proto"
	"net/http"
)

func GetError(code string, message string) []byte {
	errResponse := &tsprotos.PBError{
		Code:    code,
		Message: message,
	}
	b, _ := proto.Marshal(errResponse)
	return b
}

type JsonError struct {
	Err       error
	Code      string
	Status    int
	ExtraLogs map[string]interface{}
}

type errorMessage struct {
	Err  string `json:"err"`
	Code string `json:"code"`
}

func GetJsonError(err *JsonError, w http.ResponseWriter, r *http.Request) {
	logs := map[string]interface{}{
		"errorCode":    err.Code,
		"errorMessage": err.Err.Error(),
	}
	if err.ExtraLogs != nil && len(err.ExtraLogs) > 0 {
		for index, item := range err.ExtraLogs {
			logs[index] = item
		}
	}
	logger.LogError(logs, r)

	if w != nil {
		w.WriteHeader(err.Status)
		res := &errorMessage{
			Err:  err.Err.Error(),
			Code: err.Code,
		}
		out, _ := json.Marshal(res)
		_, _ = w.Write(out)
	}
}
