package main

import (
	"context"
	"fmt"
	"log"
	"time"
	redisgo "github.com/harrychang/redis-go/redis"
)

var ctx = context.Background()

func main() {
	fmt.Println("Go redis implementation")
	rdb := redisgo.ConnectClient()

	// Set key
	setNow := time.Now()
	redisgo.Set(ctx, "harry", 30, 100 * time.Second, rdb)
	log.Printf("Processing set redis for single command: %+v\n", time.Since(setNow))

	// Get key
	getNow := time.Now()
	val, err := redisgo.Get(ctx, "harry", rdb)
	if err != nil {
		fmt.Printf("Key harry in redis does not exists\n")
	} else {
		fmt.Printf("Key harry in redis exists with value %v\n", val)
	}
	log.Printf("Processing get redis for single command: %+v\n", time.Since(getNow))
}