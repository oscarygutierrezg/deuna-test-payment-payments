package uhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CustomSuccess(c *gin.Context, code int, msg string, data interface{}) {
	var response = Response{
		Data: data,
	}
	response.reply(c, code)
}

func Success(c *gin.Context, msg string, data interface{}) {
	CustomSuccess(c, http.StatusOK, msg, data)
}
