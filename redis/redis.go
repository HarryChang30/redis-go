package redis

import(
	"fmt"
	"context"
	"time"
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