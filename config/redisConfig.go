package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb       *redis.Client
	redisOnce sync.Once
)

func InitRedis() {

	redisOnce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			PoolSize: 10,
		})

		if _, err := rdb.Ping(context.Background()).Result(); err != nil {
			fmt.Println("redis connection error")
			rdb = nil
		} else {
			fmt.Println("redis connected succesfully")
		}

	})

}

func CloseRedis() {
	if rdb != nil {
		rdb.Close()
		fmt.Println("redis closed")
	}
}
