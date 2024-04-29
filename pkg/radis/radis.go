package radis

import (
	"github.com/hokkung/configuration/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(c config.RedisConfiguration) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})	

	return rdb, nil
}

func ProvideRedis(c config.Configuration) (*redis.Client, error) {
	return NewRedis(c.RedisConfig)
}
