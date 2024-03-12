package api

import (
	"errors"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: getEnvAccountsSID(),
	Password: getEnvAuthToken(),
})

func (app *Config) twilioSendOtp(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	res, err := client.VerifyV2.CreateVerification(getEnvServiceID(), params)

	if err != nil {
		return "", err
	}
	return *res.Sid, nil
}

func (app *Config) twilioVerifyOtp(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	fmt.Println("value of phonenumber is", phoneNumber, "value of code is", code)
	params.SetTo(phoneNumber)
	params.SetCode(code)
	res, err := client.VerifyV2.CreateVerificationCheck(getEnvServiceID(), params)
	if err != nil {
		fmt.Println("some error occured while verifying otp")
		return err
	}

	if *res.Status != "approved" {
		return errors.New("Invalid otp")
	}

	return nil

}
