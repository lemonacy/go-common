package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInt(c *gin.Context, key string) int {
	val := GetString(c, key)
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return num
}

func GetString(c *gin.Context, key string) string {
	val := c.Param(key)
	if val != "" {
		return val
	}

	val, ok := c.GetQuery(key)
	if ok {
		return val
	}

	val = c.PostForm(key)
	if val != "" {
		return val
	}

	params := map[string]interface{}{}
	if err := c.Bind(params); err != nil {
		panic(err)
	}
	v, ok := params[key]
	if ok {
		val, ok := v.(string)
		if ok {
			return val
		}
	}
	return ""
}
