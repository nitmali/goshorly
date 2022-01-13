package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var Client *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379", // use default Addr
	Password: "",           // no password set
	DB:       0,            // use default DB
})

func Init_redis() {
	_, err := Client.Ping(ctx).Result()

	if err != nil {
		log.Fatal(err.Error())
	}
}
