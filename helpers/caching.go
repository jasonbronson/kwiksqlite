package helpers

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func init() {
	// Initialize the cache
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	Cache = cache.New(5*time.Minute, 10*time.Minute)

}

func GetCache(key string) (interface{}, bool) {
	// Retrieve the value from the cache
	return Cache.Get(key)
}

func SetCache(key string, value interface{}, expiration time.Duration) {
	Cache.Set(key, value, expiration)
}
