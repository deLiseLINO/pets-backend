package login

import (
	"context"
	"net/http"
	"pets-backend/internal/handler"
	"pets-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type SignUpUserService interface {
	Add(ctx context.Context, username string, name string, email string) (*models.User, error)
	AddWithPassword(ctx context.Context, username string, name string, email string, password string) (*models.User, error)
}

type SignUpOtpService interface {
	VerifyCode(ctx context.Context, email, code string) error
}

type SignUpSSOService interface {
	HashPassword(password string) (string, error)
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

func HandleSignUp(userSvc SignUpUserService, otpSvc SignUpOtpService, ssoSvc SignUpSSOService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SignUpRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			handler.BadRequestResponse(c, err)
			return
		}

		ctx := c.Request.Context()
		err := otpSvc.VerifyCode(ctx, request.Email, request.OtpCode)
		if err != nil {
			handler.HandleVerifyCodeError(c, err)
		}

		var user *models.User

		switch request.Password {
		case "":
			user, err = userSvc.Add(ctx, request.Username, request.Name, request.Email)
			if err != nil {
				// TODO: handle duplicate errors
				handler.ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
		default:
			hashedPass, err := ssoSvc.HashPassword(request.Password)
			if err != nil {
				handler.InternalErrorResponse(c)
				return
			}

			user, err = userSvc.AddWithPassword(ctx, request.Username, request.Name, request.Email, hashedPass)
			if err != nil {
				handler.ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
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
