package recaptcha

import (
	"context"
	"errors"
	"fmt"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	recaptchaPB "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"google.golang.org/api/option"
)

type RecaptchaAssessmentRequest struct {
	ApiKey          string
	ProjectID       string
	RecaptchaKey    string
	RecaptchaToken  string
	RecaptchaAction string
	Score           float32
}

func (r *RecaptchaAssessmentRequest) Validate() (err error) {
	err = nil
	if r.ApiKey == "" {
		err = errors.New("api key is required")
	}
	if r.ProjectID == "" {
		err = errors.New("project ID is required")
	}
	if r.RecaptchaKey == "" {
		err = errors.New("recaptcha key is required")
	}
	if r.RecaptchaToken == "" {
		err = errors.New("recaptcha token is required")
	}
	if r.RecaptchaAction == "" {
		err = errors.New("recaptcha action is required")
	}
	if r.Score < 0 || r.Score > 1.0 {
		err = errors.New("recaptcha score should be between 0 and 1")
	}
	return
}

func (r *RecaptchaAssessmentRequest) ValidateRecaptcha() (response *recaptchaPB.Assessment, err error) {
	err = r.Validate()
	if err != nil {
		return
	}
	clientOption := []option.ClientOption{option.WithAPIKey(r.ApiKey)}
	// Create the reCAPTCHA client.
	ctx := context.Background()
	client, err := recaptcha.NewClient(ctx, clientOption...)
	if err != nil {
		fmt.Printf("Error creating reCAPTCHA client\n")
	}
	defer client.Close()

	// Set the properties of the event to be tracked.
	event := &recaptchaPB.Event{
		Token:   r.RecaptchaToken,
		SiteKey: r.RecaptchaKey,
	}

	assessment := &recaptchaPB.Assessment{
		Event: event,
	}

	// Build the assessment request.
	request := &recaptchaPB.CreateAssessmentRequest{
		Assessment: assessment,
		Parent:     fmt.Sprintf("projects/%s", r.ProjectID),
	}

	response, err = client.CreateAssessment(
		ctx,
		request)

	if err != nil {
		return
	}

	// Check if the token is valid.
	if !response.TokenProperties.Valid {
		err = fmt.Errorf("createAssessment() call failed because the token was invalid for the following reasons: %v",
			response.TokenProperties.InvalidReason)
		return
	}

	// Check if the expected action was executed.
	if response.TokenProperties.Action != r.RecaptchaAction {
		err = fmt.Errorf("the action attribute in your reCAPTCHA tag does not match the action you are expecting to score")
		return
	}

	// Get the risk score and the reason(s).
	// For more information on interpreting the assessment, see:
	// https://cloud.google.com/recaptcha-enterprise/docs/interpret-assessment
	// fmt.Printf("The reCAPTCHA score for this token is:  %v", response.RiskAnalysis.Score)

	// for _, reason := range response.RiskAnalysis.Reasons {
	// 	fmt.Printf(reason.String() + "\n")
	// }
	isValidScore := r.Score <= response.RiskAnalysis.Score
	if !isValidScore {
		err = errors.New("bots are not allowed to access this page")
		return
	}
	return
}
