package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseBody(c *gin.Context, x interface{}) {
	err := c.BindJSON(x)
	if err != nil {
		return
	}
}

func GetParamsId(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.Param("id"), 10, 64)
}
