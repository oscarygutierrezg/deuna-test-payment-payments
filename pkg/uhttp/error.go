package uhttp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-payments-api/internal/services"
	"payment-payments-api/pkg/util"
)

func CustomError(c *gin.Context, code int, data interface{}) {
	var response = Response{
		Data: data,
	}
	response.reply(c, code)
}

func Error(c *gin.Context, err error) {
	var status int
	var customErr *util.RequiredFieldError
	switch {
	case errors.As(err, &customErr):
		status = http.StatusBadRequest
	case errors.Is(err, services.PaymentAlreadyRefunded):
		status = http.StatusConflict
	case errors.Is(err, services.PaymentNotFound):
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}
	CustomError(c, status, err.Error())
}
