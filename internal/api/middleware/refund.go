package middleware

import (
	"github.com/gin-gonic/gin"
	"payment-payments-api/pkg/uhttp"
	"payment-payments-api/pkg/umdw"
	"payment-payments-api/pkg/util"
)

var Refund httpRefundMdw

type httpRefundMdw struct{}

func (httpRefundMdw) CreateValidation(c *gin.Context) {
	require := []string{
		"transactionId",
		"amount",
		"currency",
	}

	verify := umdw.VerificationFunctions{}

	err := umdw.BodyVerifyFields(c, require, verify)
	if err != nil {
		uhttp.Error(c, &util.RequiredFieldError{Message: err.Error()})
		return
	}

	c.Next()
}
