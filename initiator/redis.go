package initiator

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context, url, password string) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("failed to initialize redis client")
	}

	return redisClient
}
