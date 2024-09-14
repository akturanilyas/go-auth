package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Add other options as needed
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis")
	}
}
