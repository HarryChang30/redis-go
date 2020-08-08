package main

import (
	"context"
	"fmt"
	redisgo "github.com/harrychang/redis-go/redis"
	"log"
	"time"
)

var ctx = context.Background()

func main() {
	fmt.Println("Go redis implementation")
	rdb := redisgo.ConnectClient()

	// Input Scanner key and value to stored on redis
	var redisKey string
	var redisValue string
	var redisExpiry int

	fmt.Print("Enter key: ")
	fmt.Scanln(&redisKey)

	fmt.Print("Enter value: ")
	fmt.Scanln(&redisValue)

	fmt.Print("Enter expiry: ")
	fmt.Scanln(&redisExpiry)

	// Set key
	setNow := time.Now()
	redisgo.Set(ctx, redisKey, redisValue, time.Duration(redisExpiry)*time.Second, rdb)
	log.Printf("Processing set redis for single command: %+v\n", time.Since(setNow))

	// Get key
	getNow := time.Now()
	val, err := redisgo.Get(ctx, redisKey, rdb)
	if err != nil {
		fmt.Printf("Key %v in redis does not exists\n", redisKey)
	} else {
		fmt.Printf("Key %v in redis exists with value %v\n", redisKey, val)
	}
	log.Printf("Processing get redis for single command: %+v\n", time.Since(getNow))

	// Massive Single Command Set key
	var loopSize int
	fmt.Print("Enter loop size: ")
	fmt.Scanln(&loopSize)

	setMassiveNow := time.Now()
	redisgo.SingleMassiveSet(ctx, redisKey, redisValue, time.Duration(redisExpiry)*time.Second, loopSize, rdb)
	log.Printf("Processing set redis massive for single command: %+v\n", time.Since(setMassiveNow))
}
