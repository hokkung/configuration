package configuration

import (
	"context"

	ec "github.com/hokkung/configuration/entity/configuration"
	"github.com/hokkung/configuration/repository/configuration"
)

type ConfigurationService interface {
	Get(ctx context.Context, key string) (*ec.Configuration, error)
	Create(ctx context.Context, ent *ec.Configuration) error
}

type configurationService struct {
	configurationRepository configuration.ConfigurationRepository
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

	return nil
}

func NewConfigurationService(configurationRepository configuration.ConfigurationRepository) *configurationService {
	return &configurationService{
		configurationRepository: configurationRepository,
	}
}

func ProvideConfigurationService(configurationRepository configuration.ConfigurationRepository) ConfigurationService {
	return NewConfigurationService(configurationRepository)
}
