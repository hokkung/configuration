package configuration

import (
	"context"

	ec "github.com/hokkung/configuration/entity/configuration"
	"github.com/hokkung/configuration/repository/configuration"
)

type ConfigurationService interface {
	Get(ctx context.Context, key string) (*ec.Configuration, error)
	Save(ctx context.Context, key string, val string) error
}

type configurationService struct {
	configurationRepository configuration.ConfigurationRepository
}

func (c configurationService) Get(ctx context.Context, key string) (*ec.Configuration, error) {
	return nil, nil
}

func (c configurationService) Save(ctx context.Context, key string, val string) error {
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
