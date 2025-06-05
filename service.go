package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioMessage "github.com/twilio/twilio-go/rest/api/v2010"
	twilioVerify "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioVerify.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioVerify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return err
	}
	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}
	return nil
}
func (app *Config) twilioSendMSGWithService(phoneNumber string, message string) error {
	params := &twilioMessage.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetMessagingServiceSid(envMESSAGINGSERVICESID())
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	if resp.Status == nil || (*resp.Status != "queued" && *resp.Status != "sent") {
		return errors.New("error sending message")
	}
	return nil
}

