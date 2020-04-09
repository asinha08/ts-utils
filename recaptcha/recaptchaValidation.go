package recaptcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const recaptchaVerificationUrl = "https://www.google.com/recaptcha/api/siteverify"

type V3RecaptchaVerificationRequest struct {
	Secret   string  `json:"secret"`
	Response string  `json:"response"`
	Action   string  `json:"action"`
	Score    float32 `json:"score"`
	RemoteIP string  `json:"remoteip,omitempty"`
}

type V3RecaptchaVerificationResponse struct {
	Success        bool      `json:"success"`
	Score          float32   `json:"score,omitempty"`
	Action         string    `json:"action,omitempty"`
	ChallengeTs    time.Time `json:"challenge_ts"`
	Hostname       string    `json:"hostname,omitempty"`
	ErrorCodes     []string  `json:"error-codes,omitempty"`
	ApkPackageName string    `json:"apk_package_name,omitempty"`
}

func VerifyRecaptcha(req *V3RecaptchaVerificationRequest) (res *V3RecaptchaVerificationResponse, isValid bool, err error) {
	if req.Score == 0 {
		panic("recaptcha score is required")
	}
	if req.Action == "" {
		panic("recaptcha action is required")
	}

	formValues := []byte((url.Values{"secret": {req.Secret}, "response": {req.Response}}).Encode())

	response, err := http.Post(recaptchaVerificationUrl,
		"application/x-www-form-urlencoded",
		bytes.NewBuffer(formValues))
	if err != nil {
		return
	}
	defer response.Body.Close()
	recaptchaResponseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(recaptchaResponseBytes, &res)
	if err != nil {
		return
	}

	if len(res.ErrorCodes) > 0 {
		err = errors.New("invalid request : " + strings.Join(res.ErrorCodes, ", "))
		return
	}
	if res.Score <= req.Score {
		err = errors.New("bots are not allowed to access this page")
		return
	}
	if res.Action != req.Action {
		err = errors.New("invalid recaptcha action")
		return
	}
	isValid = res.Success
	return
}
