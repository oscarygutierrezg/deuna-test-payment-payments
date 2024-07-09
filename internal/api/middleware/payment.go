package middleware

import (
	"github.com/gin-gonic/gin"
	"payment-payments-api/pkg/uhttp"
	"payment-payments-api/pkg/umdw"
	"payment-payments-api/pkg/util"
)

var Payment httpPaymentMdw

type httpPaymentMdw struct{}

func (httpPaymentMdw) CreateValidation(c *gin.Context) {
	require := []string{
		"cardId",
		"cvc",
		"expiredDate",
		"amount",
		"currency",
		"merchant",
		"userId",
		"merchantId",
	}

	verify := umdw.VerificationFunctions{}

	err := umdw.BodyVerifyFields(c, require, verify)
	if err != nil {
		uhttp.Error(c, &util.RequiredFieldError{Message: err.Error()})
		return
	}

	c.Next()
}
