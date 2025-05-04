package servicecall

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/asinha08/ts-utils/tsprotos"
	"google.golang.org/protobuf/proto"
)

func Post(url string, contentType string, payload []byte) (bytesRead []byte, err error) {
	if url == "" {
		err = errors.New("bad url")
		return
	}
	if contentType == "" {
		err = errors.New("bad content type")
		return
	}
	if len(payload) == 0 {
		err = errors.New("bad payload")
		return
	}
	response, err := http.Post(url, contentType, bytes.NewBuffer(payload))
	defer func() {
		_ = response.Body.Close()
	}()
	if err != nil {
		return
	}
	bytesRead, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	statusCode := response.StatusCode
	if statusCode > 300 || statusCode < 200 {
		pbError := &tsprotos.PBError{}
		err = proto.Unmarshal(bytesRead, pbError)
		if err != nil {
			return
		}
		err = errors.New(pbError.Code + " : " + pbError.Message)
		return
	}
	return
}
