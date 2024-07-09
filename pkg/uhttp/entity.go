package uhttp

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
}

func (r Response) reply(c *gin.Context, code int) {
	if c.IsAborted() {
		return
	}
	c.JSON(code, r)
	c.Abort()
}
