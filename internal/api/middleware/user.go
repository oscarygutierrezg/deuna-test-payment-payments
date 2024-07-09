package middleware

import (
	"github.com/gin-gonic/gin"
	"payment-payments-api/pkg/uhttp"
	"payment-payments-api/pkg/umdw"
	"payment-payments-api/pkg/util"
)

var User httpUserMdw

type httpUserMdw struct{}

func (httpUserMdw) LoginValidation(c *gin.Context) {
	require := []string{
		"email",
		"password",
	}

	verify := umdw.VerificationFunctions{
		"email":    EmailValidation,
		"password": PasswordValidation,
	}

	err := umdw.BodyVerifyFields(c, require, verify)
	if err != nil {
		uhttp.Error(c, &util.RequiredFieldError{Message: err.Error()})
		return
	}

	c.Next()
}
