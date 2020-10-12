package servicecall

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func Request(method, url string, header map[string]string, payload []byte) (bytesRead []byte, err error) {
	if url == "" {
		err = errors.New("bad url")
		return
	}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	if header != nil && len(header) > 0 {
		for i, v := range header {
			request.Header.Set(i, v)
		}
	}

	timeOut := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeOut,
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}

	defer func() {
		_ = response.Body.Close()
	}()

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

	bytesRead, err = ioutil.ReadAll(response.Body)
	return
}
