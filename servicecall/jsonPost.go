package servicecall

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type TSJsonError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func JsonPost(url string, contentType string, payload []byte) (bytesRead []byte, err error) {
	if url == "" {
		err = errors.New("bad url")
		return
	}
	if contentType == "" {
		contentType = "application/json"
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
		tsError := &TSJsonError{}
		err = json.Unmarshal(bytesRead, tsError)
		if err != nil {
			return
		}
		err = errors.New(tsError.Code + " : " + tsError.Message)
		return
	}
	return
}
