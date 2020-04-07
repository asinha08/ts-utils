package recaptcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const recaptchaVerificationUrl = "https://www.google.com/recaptcha/api/siteverify"

type V3RecaptchaVerificationRequest struct {
	Secret   string  `json:"secret"`
	Response string  `json:"response"`
	Action   string  `json:"action"`
	Score    float32 `json:"score"`
}

type V3RecaptchaVerificationResponse struct {
	Success     bool      `json:"success"`
	Score       float32   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type payload struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

func VerifyRecaptcha(r *V3RecaptchaVerificationRequest) (recaptchaVerificationResponse V3RecaptchaVerificationResponse, isValid bool, err error) {
	if r.Score == 0 {
		panic("recaptcha score is required")
	}
	if r.Action == "" {
		panic("recaptcha action is required")
	}

	p := &payload{
		Secret:   r.Secret,
		Response: r.Response,
	}
	inBytes, err := json.Marshal(p)
	if err != nil {
		return
	}

	response, err := http.Post(recaptchaVerificationUrl, "application/json", bytes.NewBuffer(inBytes))
	if err != nil {
		return
	}

	recaptchaResponseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(recaptchaResponseBytes, recaptchaVerificationResponse)
	if err != nil {
		return
	}

	if len(recaptchaVerificationResponse.ErrorCodes) > 0 {
		err = errors.New("invalid request : " + strings.Join(recaptchaVerificationResponse.ErrorCodes, ", "))
		return
	}
	if recaptchaVerificationResponse.Score <= r.Score {
		err = errors.New("bots are not allowed to access this page")
		return
	}
	if recaptchaVerificationResponse.Action != r.Action {
		err = errors.New("invalid recaptcha action")
		return
	}
	isValid = recaptchaVerificationResponse.Success
	return
}
