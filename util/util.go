package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ResponseSuccess[T any](c *gin.Context, data T) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
	})
}

func ResponseErr(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"code":    1,
		"message": message,
	})
}

func ResponseErrf(c *gin.Context, template string, args ...any) {
	c.JSON(400, gin.H{
		"code":    1,
		"message": fmt.Sprintf(template, args...),
	})
}
