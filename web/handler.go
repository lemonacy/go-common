package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExceptionHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Print("System exception: %v", err)
			c.IndentedJSON(http.StatusOK, gin.H{"code": 500, "message": err})
		}
	}()

	c.Next()
}
