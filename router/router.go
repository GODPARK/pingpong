package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pingpong/config"
	"pingpong/constants"
)

func NewRouter(config *config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		if config.Pong.Token != "" {
			token := c.GetHeader(constants.PingPongTokenHeaderKey)

			if token != config.Pong.Token {
				c.Status(http.StatusUnauthorized)
				return
			}
		}

		c.Status(config.Pong.Status)
		return
	})

	return r
}
