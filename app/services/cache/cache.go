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
func Get(key string) *redis.StringCmd {
	return client.Get(key)
}

// Set a cache value
func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	client.Set("test", "qdqwdqwd", 0)
	return client.Set(key, value, expiration)
}

// Delete a cache value
func Delete(key string) *redis.IntCmd {
	return client.Del(key)
}
