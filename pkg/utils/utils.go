package utils

import (
	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context, x interface{}) {
	err := c.BindJSON(x)
	if err != nil {
		return
	}
}
