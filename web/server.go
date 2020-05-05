package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/lemonacy/go-common/config"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(ExceptionHandler)
	Router.Use(gin.Logger())

	store := memstore.NewStore([]byte("secret"))
	Router.Use(sessions.Sessions("sid", store))

	Router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}

func Run() {
	Router.Run(":" + config.Viper.GetString("server.port"))
}
