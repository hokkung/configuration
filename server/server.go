package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkung/srv/server"
)

type ServerCustomizer struct {
}

func (c *ServerCustomizer) Register(s *server.Server) {
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
