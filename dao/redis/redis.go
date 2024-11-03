package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"web2/settings"
)

var rdb *redis.Client

func Init(cfg *settings.Redisconfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: cfg.Host,

		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.Poolsize,
	})
	_, err = rdb.Ping(context.Background()).Result()
	return err
}
