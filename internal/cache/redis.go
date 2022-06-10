package cache

import (
	"os"
	"strconv"

	"github.com/millbj92/synctl/pkg/utils"
	"github.com/go-redis/redis/v8"
)

func Connect() (*redis.Client, error) {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	redisConnUrl, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}

	options := &redis.Options{
		Addr: redisConnUrl,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: dbNumber,
	}

	return redis.NewClient(options), nil
}
