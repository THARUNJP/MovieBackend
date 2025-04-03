package config

import (
	"MovieBack/internal/types"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetKey(key string, value any, expiration time.Duration) error {
	return rdb.Set(context.Background(), key, value, expiration).Err()
}

func GetKey(key string) (types.UserRefrestToken, error) {

	pipe := rdb.Pipeline()

	valCmd := pipe.Get(context.Background(), key)
	ttlCmd := pipe.TTL(context.Background(), key)

	_, err := pipe.Exec(context.Background())
	if err == redis.Nil {
		fmt.Print("No key is found")
		return types.UserRefrestToken{}, nil
	}

	value, err := valCmd.Result()
	if err == redis.Nil {
		fmt.Println("err in valcmd")
		return types.UserRefrestToken{}, nil
	}
	ttl, err := ttlCmd.Result()
	if err == redis.Nil {
		fmt.Println("err in valcmd")
		return types.UserRefrestToken{}, nil
	}
	data := types.UserRefrestToken{
		RefToken:  value,
		ExpiresAt: ttl,
	}
	return data, nil

}

func DeleteKey(key string) error {
	return rdb.Del(context.Background(), key).Err()
}
