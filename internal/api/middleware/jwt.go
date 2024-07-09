package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-payments-api/internal/models"
	"payment-payments-api/pkg/auth"
	"payment-payments-api/pkg/uhttp"
)

func JwtValidation(c *gin.Context) {
	jwtToken := c.GetHeader(auth.JwtAuthorizationHeader)

	valid, err := auth.IsJwtTokenValid(jwtToken)
	if err != nil || !valid {
		uhttp.CustomError(c, http.StatusUnauthorized, auth.JwtMessageUnauthorized)
	}

	c.Next()
}

func NewJwtToken(payload interface{}) (string, error) {

	var jwtPayload map[string]interface{}

	inBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(inBytes, &jwtPayload)
	if err != nil {
		return "", err
	}

	return auth.NewJwtToken(jwtPayload, nil)
}

func GetJwtToken(c *gin.Context) (models.User, error) {
	var user models.User
	jwtToken := c.GetHeader(auth.JwtAuthorizationHeader)

	jwtPayload, err := auth.GetJwtTokenMapClaims(jwtToken)
	if err != nil {
		return user, err
	}

	jsonBody, err := json.Marshal(jwtPayload)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(jsonBody, &user); err != nil {
		return user, err
	}

	return user, nil
}
