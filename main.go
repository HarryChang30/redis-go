package main

import (
	"context"
	"fmt"
	redisgo "github.com/harrychang/redis-go/redis"
	"log"
	"time"
)

var ctx = context.Background()

type redisConfig struct {
	key string
	value string
	expiry int
	pipeExpiry int
	loopSize int
}

func main() {
	fmt.Println("Go redis implementation")
	rdb := redisgo.ConnectClient()

	redis := redisConfig {
		"setup-key",
		"setup-value",
		3000,
		3000,
		100,
	}

	// Set key
	setNow := time.Now()
	redisgo.Set(ctx, redis.key, redis.value, time.Duration(redis.expiry)*time.Second, rdb)
	log.Printf("Processing set redis for single command: %+v\n", time.Since(setNow))

	// Get key
	getNow := time.Now()
	val, err := redisgo.Get(ctx, redis.key, rdb)
	if err != nil {
		fmt.Printf("Key %v in redis does not exists\n", redis.key)
	} else {
		fmt.Printf("Key %v in redis exists with value %v\n", redis.key, val)
	}
	log.Printf("Processing get redis for single command: %+v\n", time.Since(getNow))

	// Massive Single Command Set key
	setMassiveNow := time.Now()
	redisgo.SingleMassiveSet(ctx, redis.key, redis.value, time.Duration(redis.expiry)*time.Second, redis.loopSize, rdb)
	log.Printf("Processing set redis massive for single command: %+v\n", time.Since(setMassiveNow))

	// Pipeline single execute command line
	setPipelineNow := time.Now()
	redisgo.Pipeline(ctx, "setup-pipe-redis", time.Duration(redis.pipeExpiry)*time.Second, rdb)
	log.Printf("Processing pipeline: %+v\n", time.Since(setPipelineNow))

	// Pipeline single execute multiple command line
	setMassivePipeline := time.Now()
	n := map[string]string{
		"pipe-redis-1":  "pipe-value-1",
		"pipe-redis-2":  "pipe-value-2",
		"pipe-redis-3":   "pipe-value-3",
		"pipe-redis-4":  "pipe-value-4",
		"pipe-redis-5": "pipe-value-5",
	}
	redisgo.PipelineMassiveInsert(ctx, n, time.Duration(redis.pipeExpiry)*time.Second, rdb)
	log.Printf("Processing multiple set key value massive pipeline: %+v\n", time.Since(setMassivePipeline))
}
