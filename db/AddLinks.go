package db

import (
	"strconv"
)

func StatsIncreaseTotalLinks() (int, error) {
	val, err := Client_redis().Get(ctx, "total-links").Result()

	if err != nil {
		err := Client_redis().Set(ctx, "total-links", "0", 0).Err()
		if err != nil {
			return 0, err
		}
	}
	i, _ := strconv.Atoi(val)
	i++
	err = Client_redis().Set(ctx, "total-links", i, 0).Err()
	if err != nil {
		return 0, err
	}
	return i, nil
}
