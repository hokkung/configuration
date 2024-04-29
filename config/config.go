package config

import "github.com/kelseyhightower/envconfig"

const APP_PREFIX string = "APP"

type Configuration struct {
	RedisConfig RedisConfiguration
}

type RedisConfiguration struct {
	Addr string `envconfig:"REDIS_ADDR"`
}

func NewConfiguration() (Configuration, error) {
	var config Configuration
	err := envconfig.Process(APP_PREFIX, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func ProvideConfiguration() (Configuration, error) {
	return NewConfiguration()
}
