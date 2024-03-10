package login

import (
	"context"
	"pets-backend/internal/handler"
	"pets-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type ByCodeOtpService interface {
	VerifyCode(ctx context.Context, email string, code string) (*models.User, error)
}

type ByCodeRequest struct {
	Email   string `json:"email" validate:"required"`
	OtpCode string `json:"otp_code" validate:"required"`
}

type ByCodeResponse struct {
	Success  bool   `json:"success"`
	UserName string `json:"username"`
	Token    string `json:"token"`
}

func HandleByCode(otpSvc ByCodeOtpService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ByCodeRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			handler.BadRequestResponse(c, err)
		}

		ctx := c.Request.Context()
		user, err := otpSvc.VerifyCode(ctx, request.Email, request.OtpCode)
		if err != nil {
			// TODO: register user
			handler.InternalErrorResponse(c, err)
			return
		}

		handler.SuccessResponse(c, &ByCodeResponse{
			Success:  true,
			UserName: user.UniqueName,
		})
	}
}
