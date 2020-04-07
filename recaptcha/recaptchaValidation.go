package recaptcha

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const recaptchaVerificationUrl = "https://www.google.com/recaptcha/api/siteverify"

type V3RecaptchaVerificationRequest struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

type V3RecaptchaVerificationResponse struct {
	Success     bool      `json:"success"`
	Score       float32   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func VerifyRecaptcha(r *V3RecaptchaVerificationRequest) (recaptchaVerificationResponse V3RecaptchaVerificationResponse, err error) {
	inBytes, err := json.Marshal(r)
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

	return
}
