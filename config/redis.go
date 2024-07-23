package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

func ConnectToRedisServer() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Use the address of the Redis service
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Test the connection
	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
