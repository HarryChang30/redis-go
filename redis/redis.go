package redis

import(
	"fmt"
	"context"
	"time"
	"log"
	"github.com/go-redis/redis/v8"
)

// ConnectTest exports function
func ConnectTest() {
	fmt.Println("Connect Test package go-redis")
}

// ConnectClient exports function
func ConnectClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

// Set exports function
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration, r *redis.Client) {
	err := r.Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
}

// Get exports function
func Get(ctx context.Context, key string, r *redis.Client) (string, error){
	val, err := r.Get(ctx, key).Result()

	return val, err
}

// SingleMassiveSet exports function
func SingleMassiveSet(ctx context.Context, key string, value interface{}, expiration time.Duration, size int, r *redis.Client) {
	for i := 0; i < size; i++ {
		log.Printf("Set redis count: %v\n", i+1)
		err := r.Set(ctx, key, value, expiration).Err()
		if err != nil {
			panic(err)
		}
		log.Printf("Set redis count done: %v\n", i+1)
	}
}

// Pipeline exports function
func Pipeline(ctx context.Context, key string, expiration time.Duration, r *redis.Client) {
	pipe := r.Pipeline()

	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, expiration)

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("Pipeline executed every seconds for key: %v and value %v\n", key, incr.Val())
}

// PipelineMassiveInsert exports function
func PipelineMassiveInsert(ctx context.Context, data map[string]string, expiration time.Duration, r *redis.Client) {
	_, err := r.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for k, v := range data {
			pipe.Set(ctx, k, v, expiration).Err()
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	log.Printf("Massive Pipeline Execution is Done with multiple of commands key and value\n")
}