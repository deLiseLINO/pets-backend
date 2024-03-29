package handler

import (
	"errors"
	"net/http"
	"pets-backend/internal/models"

	"github.com/gin-gonic/gin"
)

var (
	ErrWrongCodeOrEmail = errors.New("wrong code or email")
	ErrInternalServer   = errors.New(http.StatusText(http.StatusInternalServerError))
)

type APIResponse[T any] struct {
	Data *T `json:"data,omitempty"`
}

type APIErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse[T any](c *gin.Context, data *T) {
	response := APIResponse[T]{Data: data}

	c.JSON(http.StatusOK, response)
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	apiErr := &APIError{Message: err.Error()}
	response := APIErrorResponse{Error: apiErr}

	c.AbortWithStatusJSON(statusCode, response)
}

func BadRequestResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusBadRequest, err)
}

func InternalErrorResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
}

func HandleVerifyCodeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, models.ErrMismattchCode):
		BadRequestResponse(c, ErrWrongCodeOrEmail)
		return
	case errors.Is(err, models.ErrOTPNotFound):
		BadRequestResponse(c, ErrWrongCodeOrEmail)
		return
	default:
		InternalErrorResponse(c)
		return
	}
}
