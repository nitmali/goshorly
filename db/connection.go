package db

import (
	"log"

	"github.com/go-redis/redis"
)

var Client *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

func Init_redis() {
	_, err := Client.Ping().Result()

	if err != nil {
		log.Fatal(err.Error())
	}
}
