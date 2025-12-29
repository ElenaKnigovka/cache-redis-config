package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

func NewRedisClient(ctx context.Context, addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func Get(ctx context.Context, client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}

	return val, nil
}

func Set(ctx context.Context, client *redis.Client, key string, value string, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func SetJSON(ctx context.Context, client *redis.Client, key string, value interface{}, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, jsonValue, expiration).Err()
}

func GetJSON(ctx context.Context, client *redis.Client, key string, value interface{}) error {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		return err
	}

	return json.Unmarshal([]byte(val), value)
}

func Delete(ctx context.Context, client *redis.Client, key string) error {
	return client.Del(ctx, key).Err()
}

func main() {
	ctx := context.Background()
	client, err := NewRedisClient(ctx, "localhost:6379", "", 0)
	if err != nil {
		log.Fatal(err)
	}

	key := "example_key"
	value := "example_value"
	expiration := 10 * time.Second

	if err := Set(ctx, client, key, value, expiration); err != nil {
		log.Fatal(err)
	}

	val, err := Get(ctx, client, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Value for key %s: %s\n", key, val)

	if err := Delete(ctx, client, key); err != nil {
		log.Fatal(err)
	}
}