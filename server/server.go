package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	srv "github.com/hokkung/srv/server"
)

type Customizer struct {
}

func (c *Customizer) Register(s *srv.Server) {
	s.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}

func NewCustomizer() *Customizer {
	return &Customizer{}
}

func ProvideCustomizer() (srv.ServerCustomizer, func(), error) {
	return NewCustomizer(), func() {}, nil
}
