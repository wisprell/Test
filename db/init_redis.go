package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisInstance *redis.Client

func InitRedis() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-gzhp6d9iqpnzrprcs.redis.volces.com:6379",
		Username: "root",
		Password: "Douyincloud2022", // no password set
		DB:       0,                 // use default DB
	})

	val, err := rdb.Get(ctx, "count").Result()
	if err != nil {
		err = rdb.Set(ctx, "count", "0", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("count = ", val)
	redisInstance = rdb
	return err
}

func GetRedis() *redis.Client {
	return redisInstance
}
