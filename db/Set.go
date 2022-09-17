package db

import "time"

func Set(key string, value interface{}, time time.Duration) error {
	return Client_redis().Set(ctx, key, value, time).Err()
}

func SetEX(key string, value interface{}, time time.Duration) error {
	return Client_redis().SetEX(ctx, key, value, time).Err()
}
