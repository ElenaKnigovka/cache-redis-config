package helpers

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

// GenerateRandomString generates a random string of a specified length
func GenerateRandomString(length int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var bytes = make([]byte, length)
	for i := range bytes {
		bytes[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(bytes)
}

// GetRedisClient fetches a Redis client instance or returns a default client if none is provided
func GetRedisClient(ctx context.Context, client *redis.Client) (*redis.Client, error) {
	if client != nil {
		return client, nil
	}
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}), nil
}

// HTTPErrorHandler is a custom HTTP error handler
func HTTPErrorHandler(w http.ResponseWriter, err error) {
	switch err {
	case http.ErrAbortHandler:
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "Service Unavailable")
	case http.StatusBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
	default:
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
	}
}

// SetRandomExpiryTime sets a random expiry time for a Redis key
func SetRandomExpiryTime(ctx context.Context, client *redis.Client, key string) error {
	expiry := time.Duration(rand.Intn(300)) * time.Second
	return client.Expire(ctx, key, expiry).Err()
}