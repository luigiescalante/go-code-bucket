package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

const CustomerDb = 0

type RedisManager struct {
	Db     int
	Client *redis.Client
}

func NewRedisClient(customerDb int) (*RedisManager, error) {
	address := os.Getenv("REDIS_ADDRESS")
	if address == "" {
		return nil, fmt.Errorf("REDIS_ADDRESS is not set")
	}
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("REDIS_PASSWORD is not set")
	}
	port := os.Getenv("REDIS_PORT")
	if port == " " {
		return nil, fmt.Errorf("REDIS_PORT is not set")
	}
	db := os.Getenv("REDIS_DB")
	if db == "" {
		return nil, fmt.Errorf("REDIS_DB is not set")
	}
	redisDb, err := strconv.Atoi(db)
	if err != nil {
		return nil, fmt.Errorf("REDIS_DB is not a number")
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", address, port),
		Password: password,
		DB:       redisDb,
	})
	return &RedisManager{
		Client: cli,
		Db:     customerDb,
	}, nil
}
func (rd *RedisManager) SetDb(db int) {
	rd.Db = db
}
