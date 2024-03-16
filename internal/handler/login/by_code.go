package login

import (
	"context"
	"errors"
	"pets-backend/internal/handler"
	"pets-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type ByCodeOtpService interface {
	VerifyCode(ctx context.Context, email, code string) error
}

type ByCodeUserService interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type ByCodeRequest struct {
	Email   string `json:"email" validate:"required"`
	OtpCode string `json:"otp_code" validate:"required"`
}

type ByCodeResponse struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

func HandleByCode(otpSvc ByCodeOtpService, userSvc ByCodeUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ByCodeRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			handler.BadRequestResponse(c, err)
			return
		}

		ctx := c.Request.Context()
		err := otpSvc.VerifyCode(ctx, request.Email, request.OtpCode)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrMismattchCode):
				handler.BadRequestResponse(c, handler.ErrWrongCodeOrEmail)
				return
			case errors.Is(err, models.ErrOTPNotFound):
				handler.BadRequestResponse(c, handler.ErrWrongCodeOrEmail)
				return
			default:
				handler.InternalErrorResponse(c)
				return
			}
		}
		user, err := userSvc.GetByEmail(ctx, request.Email)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrUserNotFound):
				handler.BadRequestResponse(c, models.ErrUserNotFound)
				return
			default:
				handler.InternalErrorResponse(c)
				return
			}
		}

		handler.SuccessResponse(c, &ByCodeResponse{
			UserName: user.UniqueName,
			Name:     user.Name,
		})
	}
}
