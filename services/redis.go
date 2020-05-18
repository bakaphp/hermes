package services

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

// RedisClient function for Redis connection
func RedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return client
}
