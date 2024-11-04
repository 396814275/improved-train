package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
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
func Close() {
	_ = rdb.Close()
}
