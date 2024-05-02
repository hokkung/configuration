package configuration

import (
	"context"
	"time"

	ec "github.com/hokkung/configuration/entity/configuration"
	"github.com/hokkung/configuration/pkg/id"
	"github.com/hokkung/configuration/repository/configuration"
	"github.com/hokkung/configuration/service"
)

type ConfigurationService interface {
	Get(ctx context.Context, key string) (*ec.Configuration, error)
	Create(ctx context.Context, ent *ec.Configuration) error
}

type configurationService struct {
	configurationRepository configuration.ConfigurationRepository
	entityCallbackHandler   service.EntityHandler
}

func (c configurationService) Get(ctx context.Context, key string) (*ec.Configuration, error) {
	ent, err := c.configurationRepository.FindByID(ctx, key)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

func (c configurationService) Create(ctx context.Context, ent *ec.Configuration) error {
	err := c.configurationRepository.Create(ctx, ent)
	if err != nil {
		return err
	}

	ent2 := ent
	c.entityCallbackHandler.HandleCreatedEvent(&service.CreatedEvent{
		ID:               id.RandomID(),
		EventCreatedTime: time.Now(),
		Payload:          ent2,
	})

	return nil
}

func NewConfigurationService(
	configurationRepository configuration.ConfigurationRepository,
	callbackHandler ConfigurationHandler,
) *configurationService {
	return &configurationService{
		configurationRepository: configurationRepository,
		entityCallbackHandler:   callbackHandler,
	}
}

func ProvideConfigurationService(
	configurationRepository configuration.ConfigurationRepository,
	callbackHandler ConfigurationHandler,
) ConfigurationService {
	return NewConfigurationService(
		configurationRepository,
		callbackHandler,
	)
}
