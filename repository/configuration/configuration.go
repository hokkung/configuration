package configuration

import (
	"github.com/hokkung/configuration/entity/configuration"
	"github.com/hokkung/configuration/pkg/repository"
	"github.com/hokkung/configuration/pkg/repository/radis"

	"github.com/redis/go-redis/v9"
)

type ConfigurationRepository interface {
	repository.Repository[configuration.Configuration, string]
}

type configurationRepository struct {
	*radis.BaseRedisRepository[configuration.Configuration, string]
}

func NewConfigurationRepository(rdb *redis.Client) *configurationRepository {
	return &configurationRepository{
		radis.NewBaseRedisRepository[configuration.Configuration, string](rdb),
	}
}

func ProvideConfigurationRepository(rdb *redis.Client) (ConfigurationRepository, func(), error) {
	return NewConfigurationRepository(rdb), func() {}, nil
}
