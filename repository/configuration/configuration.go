package configuration

import (
	"context"
	"encoding/json"

	"github.com/hokkung/configuration/entity/configuration"
	"github.com/redis/go-redis/v9"
)

type ConfigurationRepository interface {
	Get(ctx context.Context, key string) (configuration.Configuration, error)
	Save(ctx context.Context, config *configuration.Configuration) error
}

type configurationRepository struct {
	db *redis.Client
}

func (r *configurationRepository) get(ctx context.Context, key string) (string, error) {
	res, err := r.db.Get(ctx, key).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *configurationRepository) Get(ctx context.Context, key string) (configuration.Configuration, error) {
	res, err := r.get(ctx, key)
	if err != nil {
		return configuration.Configuration{}, err
	}

	var config configuration.Configuration
	
	err = json.Unmarshal([]byte(res), &config)
	if err != nil {
		return configuration.Configuration{}, err
	}

	return config, nil
}

func (r *configurationRepository) Save(ctx context.Context, config *configuration.Configuration) error {
	configJson, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return r.db.Set(ctx, config.GetKey(), configJson, 0).Err()
}

func NewConfigurationRepository(rdb *redis.Client) *configurationRepository {
	return &configurationRepository{
		db: rdb,
	}
}

func ProvideConfigurationRepository(rdb *redis.Client) (ConfigurationRepository, func(), error) {
	return NewConfigurationRepository(rdb), func() {}, nil
}
