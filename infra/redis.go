package infra

import (
	"api-product/config"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetRedisConnection(conf *config.RedisConfiguration) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       0, // use default DB
	})

	return rdb
}
