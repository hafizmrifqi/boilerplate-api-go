package response

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{Success: true, Data: data})
}

func Error(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{Success: false, Error: message})
}