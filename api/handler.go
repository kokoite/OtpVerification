package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranjal/golang-otp-handler/data"
)

const appTimeout = time.Second * 10

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData{PhoneNumber: payload.PhoneNumber}
		_, err := app.twilioSendOtp(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "Otp sent successfully")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()
		app.validateBody(c, &payload)
		err := app.twilioVerifyOtp(payload.User.PhoneNumber, payload.Code)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "Otp verified successfully")
	}
}
