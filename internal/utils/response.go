package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Method string
	Resp   interface{}
	Type   string
}

func SendError(c *gin.Context, status int, payload interface{}) {
	resp := Response{
		Method: c.Request.Method,
		Resp:   payload,
		Type:   "Error",
	}

	c.JSON(status, resp)
	c.Abort()
}

func SendJSON(c *gin.Context, status int, payload interface{}) {
	resp := Response{
		Method: c.Request.Method,
		Resp:   payload,
		Type:   "Success",
	}

	c.JSON(status, resp)
}
