package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Hello")
	// setup the connection to redis. "redis" is the name of container which lets
	// docker handle the networking. "mypassword" is the passwork used in docker-compose.yml
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "mypassword",
		DB:       0,
	})
	// set a key and value for testing.
	err := client.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	// Keep container alive for 3 minutes for testing
	time.Sleep(3 * time.Minute)
}
