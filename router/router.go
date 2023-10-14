package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"pingpong/config"
)

func NewRouter(config *config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		log.Println("efef")
		c.Status(config.Pong.Status)
	})

	return r
}
