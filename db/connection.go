package db

import (
	"context"
	"log"

	"git.ucode.space/Phil/goshorly/utils"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Client_redis() *redis.Client {
	var Client *redis.Client = redis.NewClient(&redis.Options{
		Addr:     utils.REDIS_URI, // use default Addr
		Password: "",              // no password set
		DB:       0,               // use default DB
	})
	return Client
}

func Init_redis() {
	_, err := Client_redis().Ping(ctx).Result()

	if err != nil {
		log.Fatal(err.Error())
	}
}
