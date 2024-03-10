package login

import (
	"context"
	"pets-backend/internal/handler"
	"pets-backend/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type GenerateOtpService interface {
	GenerateOtp() (*models.OTP, error)
	SendOtp(sendingEmail string, otpCode string) error
	SaveOtp(
		ctx context.Context,
		code string,
		email string,
		nextSendTime time.Time,
		exparationTime time.Time,
	) error
}

type GenerateOtpRequest struct {
	Email string `json:"email" validate:"required"`
}

type GenerateOtpResponse struct {
	Success        bool   `json:"success"`
	Sticker        string `json:"sticker"`
	NextSendTime   Time   `json:"next_send_time"`
	ExparationTime Time   `json:"exparation_time"`
}

type Time struct {
	TimeStamp int64  `json:"@timestamp"`
	Date      string `json:"$"`
}

func HandleGenerateOtp(otpSvc GenerateOtpService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request GenerateOtpRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			handler.BadRequestResponse(c, err)
		}

		otp, err := otpSvc.GenerateOtp()
		if err != nil {
			handler.InternalErrorResponse(c, err)
		}

		err = otpSvc.SendOtp(request.Email, otp.Code)
		if err != nil {
			handler.InternalErrorResponse(c, err)
			return
		}

		ctx := c.Request.Context()
		err = otpSvc.SaveOtp(ctx, otp.Code, request.Email, otp.NextSendTime, otp.ExparationTime)
		if err != nil {
			handler.InternalErrorResponse(c, err)
		}

		nextSendTimeResponse := Time{
			TimeStamp: otp.NextSendTime.Unix(),
			Date:      otp.NextSendTime.String(),
		}

		exparationTimeResponse := Time{
			TimeStamp: otp.ExparationTime.Unix(),
			Date:      otp.ExparationTime.String(),
		}

		handler.SuccessResponse(c, &GenerateOtpResponse{
			Success:        true,
			NextSendTime:   nextSendTimeResponse,
			ExparationTime: exparationTimeResponse,
		})
	}
}
