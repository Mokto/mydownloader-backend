package cache

import (
	"time"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// Get a cache value
func Get(key string) (string, error) {
	return client.Get(key).Result()
}

// Set a cache value
func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(key, value, expiration).Err()
}

// Delete a cache value
func Delete(key string) error {
	return client.Del(key).Err()
}

// Set a cache value
func HSet(key string, field string, value []byte) error {
	return client.HSet(key, field, value).Err()
}

// Get a cache value
func HGet(key string, field string) (string, error) {
	return client.HGet(key, field).Result()
}

// Get all values
func HGetAll(key string) (map[string]string, error) {
	return client.HGetAll(key).Result()
}

// Delete a cache value
func HDel(key string, field string) error {
	return client.HDel(key, field).Err()
}