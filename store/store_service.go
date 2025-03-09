package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define the struct wrapper around raw Redis client
// Acts as an abstraction layer around the Redis database.
// Provides a way for your application to interact with Redis without directly
// dealing with the REdis client everywhere in your code
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

//In real world usage, the cache duration shouldn't have
// an expiration time, an LRU policy config should be set where the values
// that are retrieved less often are purged automatically from
// the cache and store back in RDBMS whenever the cache is full

const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	//it initializes and mantains a connection to the Redis database
	//ensures only 1 redis client is used
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("error init Redis: %v", err))
	}

	fmt.Printf("\nReddis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}
