package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
