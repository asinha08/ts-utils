package tsError

import (
	"encoding/json"
	"net/http"

	logger "github.com/asinha08/ts-utils/tslogger/fluentd"
	"github.com/asinha08/ts-utils/tsprotos"
	"google.golang.org/protobuf/proto"
)

func GetError(code string, message string) ([]byte, error) {
	errResponse := &tsprotos.PBError{
		Code:    code,
		Message: message,
	}
	b, err := proto.Marshal(errResponse)
	return b, err
}

type JsonError struct {
	Err                error
	ClientErrorMessage string
	Code               string
	Status             int
	ExtraLogs          map[string]interface{}
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
		clientErrorMessage := err.Err.Error()
		if err.ClientErrorMessage != "" {
			clientErrorMessage = err.ClientErrorMessage
		}
		res := &errorMessage{
			Err:  clientErrorMessage,
			Code: err.Code,
		}
		out, _ := json.Marshal(res)
		_, _ = w.Write(out)
	}
}
