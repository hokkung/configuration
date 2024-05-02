package configuration

import (
	"fmt"

	"github.com/hokkung/configuration/entity/configuration"
	"github.com/hokkung/configuration/service"
)

type ConfigurationHandler interface {
	service.EntityHandler
}

type configurationHandler struct {
}

func (h configurationHandler) HandleCreatedEvent(event *service.CreatedEvent) {
	payload, ok := event.Payload.(*configuration.Configuration)
	if !ok {
		return
	}

	fmt.Printf("+%v", payload)
}

func NewConfigurationHandler() *configurationHandler {
	return &configurationHandler{}
}

func ProvideConfigurationHandler() ConfigurationHandler {
	return NewConfigurationHandler()
}
