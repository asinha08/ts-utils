package error

import (
	"github.com/asinha08/ts-utils/tsprotos"
	"github.com/golang/protobuf/proto"
)

func GetError(code string, message string) []byte {
	errResponse := &tsprotos.PBError{
		Code:    code,
		Message: message,
	}
	b, _ := proto.Marshal(errResponse)
	return b
}
