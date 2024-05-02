package radis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hokkung/configuration/pkg/repository"
	"github.com/redis/go-redis/v9"
)

const KEY_PREFIX = "config_entity"

type RedisEnity repository.Entity[string]

type BaseRedisRepository[E RedisEnity, K string] struct {
	db *redis.Client
}

func (r *BaseRedisRepository[E, K]) getKey(key string) string {
	return fmt.Sprintf("%s:%s", KEY_PREFIX, key)
}

func (r *BaseRedisRepository[E, K]) Create(ctx context.Context, ent *E) error {
	// TODO: improve this
	ent2 := *ent
	exist, err := r.FindByID(ctx, r.getKey(ent2.EntID()))
	if err != nil {
		return err
	}

	if exist != nil {
		return errors.New("key already exists")
	}

	jsonStr, err := json.Marshal(ent)
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, r.getKey(ent2.EntID()), jsonStr, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRedisRepository[E, K]) Update(ctx context.Context, ent *E) error {
	jsonStr, err := json.Marshal(ent)
	if err != nil {
		return err
	}

	// TODO: improve this
	ent2 := *ent
	err = r.db.Set(ctx, r.getKey(ent2.EntID()), jsonStr, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRedisRepository[E, K]) get(ctx context.Context, key interface{}) (string, error) {
	return r.db.Get(ctx, r.getKey(key.(string))).Result()
}

func (r *BaseRedisRepository[E, K]) FindByID(ctx context.Context, id interface{}) (*E, error) {
	res, err := r.get(ctx, id)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	var ent E
	err = json.Unmarshal([]byte(res), &ent)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (r *BaseRedisRepository[E, K]) Delete(ctx context.Context, ent E) error {
	_, err := r.db.Del(ctx, ent.EntID()).Result()
	if err != nil {
		return err
	}

	return nil
}

func NewBaseRedisRepository[E RedisEnity, K string](rdb *redis.Client) *BaseRedisRepository[E, K] {
	return &BaseRedisRepository[E, K]{
		db: rdb,
	}
}
