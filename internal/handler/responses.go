package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func InternalErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusInternalServerError, err)
}
