package db

import (
	"strconv"
)

func StatsIncreaseViewsLinks() (int, error) {
	val, err := Client_redis().Get(ctx, "total-views").Result()

	if err != nil {
		err := Client_redis().Set(ctx, "total-views", "0", 0).Err()
		if err != nil {
			return 0, err
		}
	}
	i, _ := strconv.Atoi(val)
	i++
	err = Client_redis().Set(ctx, "total-views", i, 0).Err()
	if err != nil {
		return 0, err
	}
	return i, nil
}
