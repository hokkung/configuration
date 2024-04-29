package server

import (
	"github.com/hokkung/configuration/handler/configuration"
	srv "github.com/hokkung/srv/server"
)

type Customizer struct {
	confiurationHandler *configuration.ConfigurationHandler
}

func (c *Customizer) Register(s *srv.Server) {
	configGroup := s.Engine.Group("/configurations")
	configGroup.GET("/ping", c.confiurationHandler.Ping())
	configGroup.GET("/:key", c.confiurationHandler.Get())
	configGroup.POST("", c.confiurationHandler.Create())
}

func NewCustomizer(h *configuration.ConfigurationHandler) *Customizer {
	return &Customizer{
		confiurationHandler: h,
	}
}

func ProvideCustomizer(h *configuration.ConfigurationHandler) (srv.ServerCustomizer, func(), error) {
	return NewCustomizer(h), func() {}, nil
}
