package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ExampleClient() *redis.Client {
	url := "redis://user:password@localhost:6379/0?protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}

func main() {
	fmt.Println("Golu pagla hai")
}
