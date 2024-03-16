package login

import (
	"context"
	"errors"
	"net/http"
	"pets-backend/internal/handler"
	"pets-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type SignUpUserService interface {
	Add(ctx context.Context, username string, name string, email string) (*models.User, error)
}

type SignUpOtpService interface {
	VerifyCode(ctx context.Context, email, code string) error
}

type SignUpRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	OtpCode  string `json:"otp_code"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func HandleSignUp(userSvc SignUpUserService, otpSvc SignUpOtpService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SignUpRequest
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

		// TODO: AddWithPassword
		user, err := userSvc.Add(ctx, request.Username, request.Name, request.Email)
		if err != nil {
			// TODO: handle duplicate errors
			handler.ErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		handler.SuccessResponse(c,
			&SignUpResponse{
				Username: user.UniqueName,
				Name:     user.Name,
				Email:    user.Email,
			},
		)
	}
}
