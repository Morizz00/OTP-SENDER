package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Morizz00/go-otp-sender/data"
	"github.com/gin-gonic/gin"
)

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		appTimeout := 10 * time.Second
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}
		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "OTP sent succesfully")
	}
}
func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		appTimeout := 10 * time.Second
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}
		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err:", err)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "OTP verified succesfully")
	}
}
func (app *Config) SendMSG() gin.HandlerFunc {
	return func(c *gin.Context) {
		appTimeout := 10 * time.Second
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.SendMSGData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.SendMSGData{
			PhoneNumber: payload.PhoneNumber,
			Message:     payload.Message,
		}

		err := app.twilioSendMSGWithService(newData.PhoneNumber, newData.Message)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "Message sent successfully")
	}
}
