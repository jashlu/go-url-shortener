package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
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
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nReddis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// we want to be able to save the mapping between the original Url and the generated shortUrl url
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	//we are using Go Redi library's built-in Set function
	// Stores a key/value pair in Redis
	// cacheDuration is for the expiration field, the duration after which the key will expire
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

// we should be able to retrieve the initial long URL once the short is provided
// this is when users will be calling the shortlink in the url so we need to retrieve
// the long url and redirect

func RetrieveInitialUrl(shortUrl string) string {
	// built in Get function
	// it returns a *StrignCmd object, which is a command wrapper that holds the Redis response
	// we use Result() in order to return the actual value and any possible error
	originalUrl, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return originalUrl
}
